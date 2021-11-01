package controller

import (
	"log"
	"net/http"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// get parameters
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// data validation
	// gin.H == map[string]interface{}
	// log.Println(name, telephone, password)
	if len(telephone) != 11 {
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
	// if there is no name, pass a random 10 bit string.
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	//determine whether the mobile-phone nember exists
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "user is exist",
		})
		return
	}

	//if user is not exist ,then create user
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// return results
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
