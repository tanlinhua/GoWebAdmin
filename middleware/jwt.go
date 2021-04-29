package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/pkg/config"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

var JwtKey = []byte(config.JwtKey)

type Claims struct {
	UserName string
	Id       int
	jwt.StandardClaims
}

// 生成token
func GetJWT(username string, id int) (bool, string) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName: username,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //生成时间
			Issuer:    "127.0.0.1",       // 签名颁发者
			Subject:   "userToken",       //签名主题
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return false, ""
	}
	return true, token
}

// jwt中间件
func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if len(tokenString) == 0 {
			response.New(c).Error(-1, "TOKEN不存在")
			c.Abort()
			return
		}

		token, claims, err := parseToken(tokenString)
		if err != nil || !token.Valid {
			response.New(c).Error(-2, "TOKEN错误")
			c.Abort()
			return
		}

		fmt.Println(config.AdminName, claims.UserName, claims.Id)

		if config.AdminName != claims.UserName {
			ok, msg := CheckPermission(claims.UserName, c.Request.RequestURI, c.Request.Method)

			fmt.Println(ok, msg)

			if !ok {
				response.New(c).Error(-3, msg)
				c.Abort()
			}
		}
		c.Set("id", claims.Id)
		c.Next()
	}
}

// 解析token
func parseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, Claims, err
}

// 验证用户是否有该uri操作权限
func CheckPermission(username, uri, method string) (bool, string) {
	// return model.CheckPermission(username, uri, method)
	return false, "权限不足"
}
