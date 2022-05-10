package middleware

import (
	"fmt"
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(utils.JwtKey)

// var jwtKey = []byte("a_secret_crect")

type MyClaims struct {
	// Username string `json:"username"`
	// Password string `json:"password"`
	UserId uint
	jwt.StandardClaims
}

// 生成token
func SetToken(user model.User) (string, error) {
	expireTime := time.Now().Add(1 * 24 * time.Hour)
	myClaims := MyClaims{
		// Username: user.Name,
		// Password: user.Password,
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "blackunicorn",
			Subject:   "user token",
		},
	}
	// 生成token结构体
	// {"alg":"HS256","typ":"JWT"} . {"UserId":3,"exp":1583418059,"iat":1582813259,"iss":"oceanlearn.tech","sub":"user token"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 验证token
func CheckToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	myClaims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, myClaims, err
}

// jwt中间件
func TokenJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token
		tokenString := ctx.Request.Header.Get("Authorization")
		fmt.Println("tokenString")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions!",
			})
			ctx.Abort()
			return
		}

		// 校验token
		// checkTokenString := strings.SplitN(tokenString, " ", 2)
		// 另一种方式
		tokenString = tokenString[7:]
		token, claims, err := CheckToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions!!",
			})
			ctx.Abort()
			return
		}
		// token过期
		if time.Now().Unix() > claims.ExpiresAt {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token is over!",
			})
		}

		// 验证token成功后， 获取用户信息
		userId := claims.UserId
		DB := model.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 检验用户信息
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions!!!",
			})
		}

		// 如果用户存在，把用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}
}
