package controller

import (
	"log"
	"strconv"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost response.CreatePostRequest
	// data validation
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "data validation error1")
		return
	}

	// get  user info
	user, _ := ctx.Get("user")
	// create post
	post := model.Post{
		// UserId: 1,
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)

	}
	response.Success(ctx, gin.H{"post": post}, "post created success")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost response.CreatePostRequest
	// data validation
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "data validation error1")
		return
	}
	//get id in path
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).RowsAffected < 1 {
		response.Fail(ctx, nil, "post does not exist!")
		return
	}

	// determine whether the current user is the author of the post
	// get  user info
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "The post does not belong to you. Please do not operate illegally.")
		return
	}
	// update post
	// Note that gorm2 needs to be used "Updates" s s s. update multiple attributes.
	if err := p.DB.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(ctx, nil, "post updated fail.")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "post updated success")

}

func (p PostController) Show(ctx *gin.Context) {
	//get id in path
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Preload("Category").Where("id=?", postId).First(&post).RowsAffected < 1 {
		response.Fail(ctx, nil, "post does not exist!")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "post show success")
}

func (p PostController) Delete(ctx *gin.Context) {
	//get id in path
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).RowsAffected < 1 {
		response.Fail(ctx, nil, "post does not exist!")
		return
	}
	// determine whether the current user is the author of the post
	// get  user info
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "The post does not belong to you. Please do not operate illegally.")
		return
	}
	p.DB.Delete(&post)
	response.Success(ctx, gin.H{"post": post}, "post delete success")
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "1"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 查询总条数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)
	response.Success(ctx, gin.H{"data": gin.H{"posts": posts, "total": total}}, "success")
}
