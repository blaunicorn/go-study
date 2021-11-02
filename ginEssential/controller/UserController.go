package controller

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/dto"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/response"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 第一版，没有加入自定义状态码
// func Register(ctx *gin.Context) {
// 	DB := common.GetDB()

// 	// get parameters
// 	name := ctx.PostForm("name")
// 	telephone := ctx.PostForm("telephone")
// 	password := ctx.PostForm("password")

// 	// data validation
// 	// gin.H == map[string]interface{}
// 	// log.Println(name, telephone, password)
// 	if len(telephone) != 11 {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "Mobile phone number must be 11 digits",
// 		})
// 		return
// 	}
// 	if len(password) < 6 {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "The password cannot be less than 6 digits",
// 		})
// 		return
// 	}
// 	// if there is no name, pass a random 10 bit string.
// 	if len(name) == 0 {
// 		name = util.RandomString(10)
// 	}
// 	log.Println(name, telephone, password)

// 	//determine whether the mobile-phone nember exists
// 	if isTelephoneExist(DB, telephone) {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "user is exist",
// 		})
// 		return
// 	}

// 	//if user is not exist ,then create user
// 	// Encrypt user password
// 	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"code": 500,
// 			"msg":  "Encrypt is error",
// 		})
// 		return
// 	}
// 	newUser := model.User{
// 		Name:      name,
// 		Telephone: telephone,
// 		Password:  string(hasedPassword),
// 	}
// 	DB.Create(&newUser)

// 	// return results
// 	ctx.JSON(200, gin.H{
// 		"message": "success",
// 	})
// }

// func Login(ctx *gin.Context) {
// 	DB := common.GetDB()
// 	// get parameters
// 	telephone := ctx.PostForm("telephone")
// 	password := ctx.PostForm("password")
// 	fmt.Println("form:", telephone, password)
// 	// data validation
// 	if len(telephone) != 11 {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "Mobile phone number must be 11 digits",
// 		})
// 		return
// 	}
// 	if len(password) < 6 {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "The password cannot be less than 6 digits",
// 		})
// 		return
// 	}

// 	// determine whether the mobile-phone nember exists
// 	var user model.User
// 	DB.Where("telephone=?", telephone).First(&user)
// 	if user.ID == 0 {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"code": 422,
// 			"msg":  "user is not exist",
// 		})
// 		return
// 	}

// 	// Determine whether the password is correct
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"code": 400,
// 			"msg":  "The password error",
// 		})
// 		return
// 	}

// 	//  issue token
// 	// token := "11"
// 	token, err := common.ReleaseToken(user)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"code": 500,
// 			"msg":  "System exception",
// 		})
// 		log.Printf("token generate error : %v", err)
// 		return
// 	}

// 	// return results
// 	ctx.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "success",
// 		"data":    gin.H{"token": token},
// 	})
// }

// func Info(ctx *gin.Context) {
// 	user, _ := ctx.Get("user")
// 	fmt.Println(user)
// 	fmt.Println("type:", reflect.TypeOf(user))
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": 200,
// 		"msg":  "success",
// 		// "data": gin.H{"user": user},  // bug. show password
// 		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
// 	})
// }

// 第二版 ，增加状态码
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
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "Mobile phone number must be 11 digits",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Mobile phone number must be 11 digits")
		return
	}
	if len(password) < 6 {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "The password cannot be less than 6 digits",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "The password cannot be less than 6 digits")
		return
	}
	// if there is no name, pass a random 10 bit string.
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	//determine whether the mobile-phone nember exists
	if isTelephoneExist(DB, telephone) {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"code": 422,
		// 	"msg":  "user is exist",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "user is exist")
		return
	}

	//if user is not exist ,then create user
	// Encrypt user password
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{
		// 	"code": 500,
		// 	"msg":  "Encrypt is error",
		// })
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "Encrypt is error")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// return results
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// get parameters
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	fmt.Println("form:", telephone, password)
	// data validation
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

	// determine whether the mobile-phone nember exists
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "user is not exist",
		})
		return
	}

	// Determine whether the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "The password error",
		})
		return
	}

	//  issue token
	// token := "11"
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "System exception",
		})
		log.Printf("token generate error : %v", err)
		return
	}

	// return results
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    gin.H{"token": token},
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	fmt.Println(user)
	fmt.Println("type:", reflect.TypeOf(user))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		// "data": gin.H{"user": user},  // bug. show password
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
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
