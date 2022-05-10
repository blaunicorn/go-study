package v1

import (
	"encoding/base64"
	"fmt"
	"gin-vue-lifeassistant/middleware"
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/responsemsg"
	"gin-vue-lifeassistant/utils/validator"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type IUserController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	List(ctx *gin.Context)
}

type UserController struct {
	DB *gorm.DB
}

func NewUserController() IUserController {
	db := model.GetDB()
	db.AutoMigrate(model.User{})
	return UserController{DB: db}
	// 18节修改
	// repository := repository.NewCategoryRepository()
	// repository.DB.AutoMigrate(model.Category{})
	// return CategoryController{Repository: repository}
}

// 添加用户
func (p UserController) Create(ctx *gin.Context) {
	fmt.Println("------")
	// 从原有Request.Body读取
	// body, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(body))
	var requestUser = model.User{}
	ctx.ShouldBindJSON(&requestUser)
	fmt.Println(requestUser)

	// 新增校验功能
	var msg string
	var code int
	msg, code = validator.Validate(&requestUser)
	if code != 200 {
		responsemsg.Response(ctx, http.StatusUnprocessableEntity, 422, nil, msg)
		return
	}

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	role := requestUser.Role

	if isAttributeValueExist(p.DB, "telephone", telephone) {
		responsemsg.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "user is exist")
		return
	}
	//if user is not exist ,then create user
	// mothod one： Encrypt user password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	// ctx.JSON(http.StatusInternalServerError, gin.H{
	// 	// 	"code": 500,
	// 	// 	"msg":  "Encrypt is error",
	// 	// })
	// 	responsemsg.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "Encrypt is error")
	// 	return
	// }

	// mothod two ScryptPwd
	hashedPassword := ScryptPwd(password)
	// write to basedb
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
		Role:      role,
		IsDel:     0,
	}
	p.DB.Create(&newUser)
	responsemsg.Success(ctx, gin.H{"user": &newUser}, "user created success")
}

// 查询单个用户

func (p UserController) Show(ctx *gin.Context) {
	//get id in path
	// id := ctx.Params.ByName("id")

}

// 查询用户列表
func (p UserController) List(ctx *gin.Context) {

	// 链式不定参数查询
	DB := p.DB
	// if name, isExist := ctx.GetQuery("username"); isExist == true {
	// 	//模糊查询
	// 	DB = DB.Where("name like ?", "%"+name+"%")
	// }
	// if telephone, isExist := ctx.GetQuery("telephone"); isExist == true {

	// 	DB = DB.Where("telephone = ?", telephone)
	// }

	// 定义一个map[查询参数] 查询条件 然后根据map 查询，map也可以放在sql数据库中
	mapValue := map[string]string{
		"name":      "like",
		"telephone": "=",
		"age":       ">",
	}
	for key, value := range mapValue {
		if queryValue, isExist := ctx.GetQuery(key); isExist {
			DB = DB.Where(key+" "+value+" ?", queryValue)
		}
	}
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	// desc：降序排列  asc 升序 默认
	sort := ctx.DefaultQuery("sort", "desc")

	// var users []model.User
	// 新建接口，去掉password字段，避免密码泄露
	type User struct {
		gorm.Model
		ID        uint
		Name      string
		RealName  string
		Telephone string
		Role      int
	}
	var users []User
	var total int64
	DB = DB.Order("created_at " + sort).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users).Count(&total)

	// 查询总条数
	// var total int64
	// DB.Model(model.User{}).Count(&total)

	responsemsg.Success(ctx, gin.H{"data": gin.H{"users": users, "total": total}}, "success")
}

// 编辑修改用户
func (p UserController) Update(ctx *gin.Context) {
	// 忽略密码项
	// 使用update时，stuct 存在0值不会更新的问题，所以用map更新
	var data model.User
	ctx.ShouldBindJSON(&data)
	var user = make(map[string]interface {
	})
	user["ID"] = ctx.Param("id")
	user["Name"] = data.Name
	user["RealName"] = data.RealName
	user["Telephone"] = data.Telephone
	user["Role"] = data.Role
	telephone := data.Telephone

	//get id in path
	Id := ctx.Params.ByName("id")

	if p.DB.Where("id=? and is_del=0", Id).First(&data).RowsAffected < 1 {
		responsemsg.Fail(ctx, nil, "user ID does not exist!")
		return
	}

	if isAttributeValueExist(p.DB, "telephone", telephone) {
		responsemsg.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "user 存在手机号")
		return
	}
	// determine whether the current user is the author of the post
	// get  user info
	// user, _ := ctx.Get("user")
	// userId := user.(model.User).ID
	// if userId != post.UserId {
	// 	response.Fail(ctx, nil, "The post does not belong to you. Please do not operate illegally.")
	// 	return
	// }
	// update post
	// Note that gorm2 needs to be used "Updates" s s s. update multiple attributes.
	if err := p.DB.Where("id= ? and is_del=0", Id).Updates(user).Error; err != nil {
		responsemsg.Fail(ctx, nil, "user updated fail.")
		return
	}
	responsemsg.Success(ctx, gin.H{"user": user}, "user updated success")

}

// 删除用户
func (p UserController) Delete(ctx *gin.Context) {
	//get id in path
	id := ctx.Params.ByName("id")
	var user model.User
	if p.DB.Where("id=?", id).First(&user).RowsAffected < 1 {
		responsemsg.Fail(ctx, nil, "用户id不存在！")
		return
	}
	// determine whether the current user is the author of the post
	// get  user info,通过token获取当前用户信息，并判断是否有删除权限
	// user, _ := ctx.Get("user")
	// userId := user.(model.User).ID
	// if userId != post.UserId {
	// 	response.Fail(ctx, nil, "The post does not belong to you. Please do not operate illegally.")
	// 	return
	// }
	user.IsDel = 1
	p.DB.Where("id=?", id).Updates(&user)
	// p.DB.Where("id=?", id).Delete(&user)
	responsemsg.Success(ctx, gin.H{"userId": id}, "user delete success")

}

// 删除多个用户
func (p UserController) DeleteList(ctx *gin.Context) {

}

// 用户登录 生成token
func Login(ctx *gin.Context) {
	DB := model.GetDB()
	var data = model.User{}
	ctx.ShouldBindJSON(&data)
	fmt.Println(&data)

	//get Json parameters
	telephone := data.Telephone
	password := data.Password

	// get Form parameters
	// telephone := ctx.PostForm("telephone")
	// password := ctx.PostForm("password")
	// fmt.Println("form:", telephone, password)

	// data validation
	if len(telephone) != 11 && len(telephone) != 12 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "Mobile phone number must be 11 digits",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The password cannot be less than 6 digits",
		})
		return
	}
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	// determine whether the mobile-phone nember exists
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "user is not exist",
		})
		return
	}

	// Determine whether the password is correct
	if ScryptPwd(password) != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "password erro!",
		})
		return
	}

	// issue token
	token, err := middleware.SetToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "System exception",
		})
		log.Printf("token generate error : %v", err)
		return
	}

	// return results
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}

// 查询用户、电话等属性的值是否在数据库中存在
func isAttributeValueExist(db *gorm.DB, AttributeName string, AttributeValue string) bool {
	var user model.User
	db.Where(AttributeName+"=? and is_del = 0 ", AttributeValue).First((&user))
	fmt.Println(user, AttributeName, AttributeValue)
	return user.ID != 0
}

// 密码加密
func ScryptPwd(password string) string {

	const KeyLen = 18
	salt := make([]byte, 8)
	salt = []byte{12, 23, 40, 23, 45, 64, 45, 67}
	passwordByScrypt, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	newPassword := base64.StdEncoding.EncodeToString(passwordByScrypt)
	return newPassword
}
