package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	var pool = GetRedis("redis://81.70.77.41:6379")
	conn := pool.Get()
	defer conn.Close()
	key := "test_redis_key"
	key1 := "test_redis_key11"
	value := "test_redis_value11"
	_, e := conn.Do("SETEX", key, 60*60*24, value)
	if e != nil {
		panic(e)
	}
	getValue, e := redis.String(conn.Do("GET", key1))
	fmt.Println("receive from redis key 'test_redis_key':", getValue)
}
func GetRedis(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     200,               //最大空闲数
		MaxActive:   100,               //最大活跃数
		IdleTimeout: 180 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.DialURL(url)
			// conn, err := redis.Dial("tcp", "81.70.77.41:6379")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			//此处1234对应redis密码
			if _, err := conn.Do("AUTH", "123456"); err != nil {
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
