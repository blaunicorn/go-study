package v1

// 添加文章
// 查询单个文章
// 查询文章列表
// 编辑修改文章
// 删除文章

import (
	"fmt"
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/responsemsg"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IArticleController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteList(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	List(ctx *gin.Context)
}

type ArticleController struct {
	DB *gorm.DB
}

func NewArticleController() IArticleController {
	db := model.GetDB()
	db.AutoMigrate(model.Article{})
	return ArticleController{DB: db}
}

// 添加文章，不需要验证
func (p ArticleController) Create(ctx *gin.Context) {
	fmt.Println("Article:")
	// 从原有Request.Body读取
	// body的内容只能读取一次，后面在读取都是空的。
	// 想在中间件中获取body请求参数，记录到日志。如果body只能获取一次
	// 在c.Request.Body.Read()之后回设body：c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	// body, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(body))

	var data = model.Article{}
	// 方式一
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// 方式二
	// requestJson := make(map[string]interface{})
	// 方式三
	ctx.BindJSON(&data)
	// // requestJson.Name = "aa"
	fmt.Println(&data)

	// write to basedb
	data.IsDel = 0
	p.DB.Create(&data)
	responsemsg.Success(ctx, gin.H{"Article": &data}, "Article created success")
}

// 查询单个文章

func (p ArticleController) Show(ctx *gin.Context) {
	//get id in path
	id := ctx.Params.ByName("id")
	data := model.Article{}
	if p.DB.Where("id=? and is_del=0", id).Preload("Category").First(&data).RowsAffected < 1 {
		responsemsg.Fail(ctx, nil, "article ID does not exist!")
		return
	}
	responsemsg.Success(ctx, gin.H{"Article": &data}, "Article find success")
}

// 查询文章列表
func (p ArticleController) List(ctx *gin.Context) {

	// select * from article inner join category on article.category_id = category.id

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
		"title":       "like",
		"category_id": "=",
		"age":         ">",
	}
	for key, value := range mapValue {
		if queryValue, isExist := ctx.GetQuery(key); isExist {
			fmt.Println(key, value, queryValue)
			DB = DB.Where(key+" "+value+" ? and is_del=0", queryValue)
		}
	}
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	// desc：降序排列  asc 升序 默认
	sort := ctx.DefaultQuery("sort", "desc")

	var articleList []model.Article

	DB = DB.Preload("Category").Order("created_at " + sort).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articleList)

	// 查询总条数
	var total int64
	DB.Model(model.Article{}).Count(&total)

	responsemsg.Success(ctx, gin.H{"data": gin.H{"articles": articleList, "total": total}}, "success")
}

// 编辑修改文章
func (p ArticleController) Update(ctx *gin.Context) {
	// 忽略密码项
	// 使用update时，stuct 存在0值不会更新的问题，所以用map更新
	var data model.Article
	ctx.ShouldBindJSON(&data)
	var article = make(map[string]interface {
	})
	article["ID"] = ctx.Param("id")
	article["UserId"] = data.UserId
	article["CategoryId"] = data.CategoryId
	article["Title"] = data.Title
	article["HeadImg"] = data.HeadImg
	article["Desc"] = data.Desc
	article["Content"] = data.Content

	//get id in path
	Id := ctx.Params.ByName("id")

	if p.DB.Where("id=? and is_del=0", Id).First(&data).RowsAffected < 1 {
		responsemsg.Fail(ctx, nil, "article ID does not exist!")
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
	if err := p.DB.Model(&data).Where("id= ? and is_del=0", Id).Updates(&article).Error; err != nil {
		fmt.Println(err)
		responsemsg.Fail(ctx, nil, "data updated fail.")
		return
	}
	responsemsg.Success(ctx, gin.H{"article": &article}, "article updated success")

}

// 删除文章
func (p ArticleController) Delete(ctx *gin.Context) {
	//get id in path
	id := ctx.Params.ByName("id")
	var article model.Article
	if p.DB.Where("id=? and is_del=0", id).First(&article).RowsAffected < 1 {
		responsemsg.Fail(ctx, nil, "文章id不存在！")
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
	article.IsDel = 1
	p.DB.Where("id=? and is_del=0", id).Updates(&article)
	// p.DB.Where("id=?", id).Delete(&user)
	responsemsg.Success(ctx, gin.H{"articleId": id}, "article delete success")

}

// 删除多个文章
func (p ArticleController) DeleteList(ctx *gin.Context) {

}

// 查询文章属性的值是否在数据库中存在
func isAttributeValueExistInArticle(db *gorm.DB, AttributeName string, AttributeValue string) bool {
	var article model.Article
	db.Where(AttributeName+"=? and is_del=0", AttributeValue).First(&article)
	return article.ID != uuid.Nil
}
