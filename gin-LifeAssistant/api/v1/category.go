package v1

import (
	"fmt"
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/responsemsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查询分类是否存在
// 添加分类
// 根据id查询单个分类
// 根据id查询单个分类下的文章
// 查询分类列表
// 编辑修改分类
// 删除分类
type ICategoryController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := model.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

// 添加分类
func (p CategoryController) Create(ctx *gin.Context) {
	fmt.Println("------")
	// 从原有Request.Body读取
	// body的内容只能读取一次，后面在读取都是空的。
	// 想在中间件中获取body请求参数，记录到日志。如果body只能获取一次
	// 在c.Request.Body.Read()之后回设body：c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	// body, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(body))

	var requestJson = model.Category{}
	// 方式一
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// 方式二
	// requestJson := make(map[string]interface{})
	// 方式三
	ctx.BindJSON(&requestJson)
	// // requestJson.Name = "aa"
	fmt.Println(&requestJson)
	if isAttributeValueExistInCategory(p.DB, "name", requestJson.Name) {
		responsemsg.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "category name is exist")
		return
	}

	// write to basedb

	p.DB.Create(&requestJson)
	responsemsg.Success(ctx, gin.H{"category": &requestJson}, "category created success")
}

// 查询单个分类

func (p CategoryController) Show(ctx *gin.Context) {
	//get id in path
	// id := ctx.Params.ByName("id")

}

// 查询分类列表
func (p CategoryController) List(ctx *gin.Context) {

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
		if queryValue, isExist := ctx.GetQuery(key); isExist == true {
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
	DB = DB.Order("created_at " + sort).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users)

	// 查询总条数
	var total int64
	DB.Model(model.User{}).Count(&total)

	responsemsg.Success(ctx, gin.H{"data": gin.H{"users": users, "total": total}}, "success")
}

// 编辑修改分类
func (p CategoryController) Update(ctx *gin.Context) {
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

// 删除分类
func (p CategoryController) Delete(ctx *gin.Context) {
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

// 删除多个分类
func (p CategoryController) DeleteList(ctx *gin.Context) {

}

// 查询分类属性的值是否在数据库中存在
func isAttributeValueExistInCategory(db *gorm.DB, AttributeName string, AttributeValue string) bool {
	var category model.Category
	db.Where(AttributeName+"=? ", AttributeValue).First(&category)
	return category.ID != 0
}
