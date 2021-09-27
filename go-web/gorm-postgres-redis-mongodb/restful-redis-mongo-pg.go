package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var pool = GetRedis("redis://81.70.77.41:6379")

func GetRedis(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     200,               //最大空闲数
		MaxActive:   100,               //最大活跃数
		IdleTimeout: 180 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			// conn, err := redis.DialURL(url)
			conn, err := redis.Dial("tcp", "81.70.77.41:6379")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			//此处1234对应redis密码
			if _, err := conn.Do("AUTH", "123456"); err != nil {
				fmt.Println(err)
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (u *User) SyncRedis(conn redis.Conn) {
	if conn == nil {
		conn = pool.Get()
		defer conn.Close()
	}
	buf, _ := json.Marshal(u)
	key := fmt.Sprintf("test:user_info:%d", u.Id)
	_, e := conn.Do("SETEX", key, 60*60*60, buf)
	if e != nil {
		panic(e)
	}

}

var db *gorm.DB
var mogSession *mgo.Session
var col *mgo.Collection

func init() {
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			"81.70.77.41",
			"8030",
			"gitchat",
			"test",
			"disable",
			"123456",
		))
	db.SingularTable(true) //设置全局表名禁用复数
	db.LogMode(true)       // 启用Logger，显示详细日志
	db.DB().SetConnMaxIdleTime(10 * time.Second)
	db.DB().SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	}

	mogSession := InitMongoSession()
	col = mogSession.DB("testdb").C("user_info")
	if err != nil {
		fmt.Println("CollectionNames-error:", err)
		panic(err)
	}
	fmt.Println(col)
}
func clear() {
	mogSession.Close()
	db.Close()
}

func InitMongoSession() *mgo.Session {
	mUsername := "admin" //mongodb的账号
	mPassword := "admin" //mongodb的密码
	mgoSession, err := mgo.Dial("81.70.104.105:27017")
	if err != nil {
		fmt.Println("mgo.Dial-error:", err)
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	mgoSession.SetMode(mgo.Monotonic, true)
	myDB := mgoSession.DB("admin") //这里的关键是连接mongodb后，选择admin数据库，然后登录，确保账号密码无误之后，该连接就一直能用了
	//出现server returned error on SASL authentication step: Authentication failed. 这个错也是因为没有在admin数据库下登录
	err = myDB.Login(mUsername, mPassword)
	if err != nil {
		fmt.Println("Login-error:", err)
		panic(err)
	}

	// mogSession.DB("testdb").C("user_info")
	mgoSession.SetPoolLimit(10)
	return mgoSession
}

type VisitLog struct {
	URL         string    `json:"url"`
	IP          string    `json:"ip"`
	ContentType string    `json:"content-type"`
	Body        []byte    `json:"body"`
	Query       string    `json:"query"`
	CreatedAt   time.Time `json:"create_at"`
}

func VisitLogMiddleware(c *gin.Context) {
	defer c.Next()
	var vl VisitLog
	vl.URL = c.Request.URL.String()
	vl.ContentType = c.ContentType()
	vl.IP = c.ClientIP()
	buf, _ := ioutil.ReadAll(c.Request.Body)
	if len(buf) != 0 {
		vl.Body = buf
		fmt.Println(string(buf[0 : len(buf)-0]))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf))

	}
	vl.Query = c.Request.URL.Query().Encode()
	vl.CreatedAt = time.Now()
	fmt.Println(&vl)
	err := col.Insert(&vl)
	if err != nil {
		fmt.Println("01mongo error:", err.Error())
		panic(err)
	}
}
func main() {
	defer clear()
	r := gin.Default()
	r.Use(VisitLogMiddleware)
	r.GET("/user/", get)
	r.GET("user/:id/", getOne)
	r.POST("/user/", post)
	r.PATCH("/user/:id/", patch)
	r.DELETE("/user/:id/", deleteById)
	r.GET("/visit-log/", visitLogControl)
	r.Run(":8082")
}

func get(c *gin.Context) {
	var users []User
	if e := db.Raw("select * from user_info").Scan(&users).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	conn := pool.Get()
	defer conn.Close()
	for i, _ := range users {
		users[i].SyncRedis(conn)
	}
	c.JSON(200, gin.H{"messge": "success", "data": users})
}

func getOne(c *gin.Context) {
	var user User
	id := c.Param("id")
	conn := pool.Get()
	defer conn.Close()

	buf, err := redis.Bytes(conn.Do("GET", fmt.Sprintf("gitchat:user_info:%s", id)))
	if err != nil {
		fmt.Println("1", buf)
		c.JSON(500, gin.H{"message": err.Error()})
		panic(err)
	}
	fmt.Println("2", buf)
	if len(buf) != 0 {
		err = json.Unmarshal(buf, &user)
		if err != nil {
			panic(err)
		}
		fmt.Println("3", buf, user)
	} else {
		fmt.Println("4", buf, &user)
		db.Raw("select * from user_info where id=?", id).Scan(&user)
	}
	c.JSON(200, gin.H{"message": "success", "data": user})
}

func post(c *gin.Context) {
	var user User
	if e := c.Bind(&user); e != nil {

		c.JSON(500, gin.H{"message": e.Error()})
		panic(e)
	}
	if e := db.Raw("insert into user_info(username) values(?) returning *", user.Username).Scan(&user).Error; e != nil {
		c.JSON(500, gin.H{"message": e.Error()})
		return
	}
	user.SyncRedis(nil)
	c.JSON(200, gin.H{"message": "success", "data": user})

}
func patch(c *gin.Context) {
	var user User
	c.Bind(&user)
	id := c.Param("id")
	if id == "" {
		c.JSON(200, gin.H{"message": "id值不能为空"})
		return
	}
	db.Raw("update user_info set username=? where id=? returning *", user.Username, id).Scan(&user)
	user.SyncRedis(nil)
	c.JSON(200, gin.H{"message": "success", "data": user})
}
func deleteById(c *gin.Context) {
	id := c.Param("id")
	row := db.Exec("delte from user_info where id=?", id)
	fmt.Println(row)
	c.JSON(200, gin.H{"message": "success"})
}

// mogondb 中间件
func visitLogControl(c *gin.Context) {
	var results []VisitLog
	err := col.Find(nil).All(&results)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		panic(err)
	}
	c.JSON(200, gin.H{"message": "sucess", "results": results})
}
