package controller

import (
	"strconv"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/repository"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/response"
	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	// 18节前
	// DB *gorm.DB

	// 18节修改
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	// 18节前
	// db := common.GetDB()

	// db.AutoMigrate(model.Category{})

	// return CategoryController{DB: db}

	// 18节修改
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {

	var requestCategory response.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "Data validation error, category name is required!")
		return
	}
	// var requestCategory model.Category
	// 	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "Data validation error, category name is required")
		return
	}

	// c.DB.Create(&requestCategory)

	// 18节前
	// categroy := model.Category{Name: requestCategory.Name}
	// c.DB.Create(&categroy)
	// response.Success(ctx, gin.H{"category": requestCategory}, "success")

	// 18节修改开始
	// var category *model.Category
	// var err error
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "创建失败")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "success")
	// 18节修改 结束

}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "Data validation error, category name is required")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")
		return
	}

	// 更新分类
	// 方式1 map
	// 方式2 结构体
	// 方式3 name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "success")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "success")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")
		return
	}
	if err := c.DB.Delete(&model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{"err": err}, "删除失败")
		return
	}
	response.Success(ctx, nil, "delete success")
}
