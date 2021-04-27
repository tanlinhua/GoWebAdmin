package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/config"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

var JwtKey = []byte(config.JwtKey)

type Claims struct {
	UserName string
	Flag     string
	jwt.StandardClaims
}

// 生成token
func GetJWT(username, flag string) (bool, string) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName: username,
		Flag:     flag,
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
			c.JSON(http.StatusOK, response.Error("TOKEN不存在", 0, nil))
			c.Abort()
			return
		}

		token, claims, err := parseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, response.Error("TOKEN错误", 0, nil))
			c.Abort()
			return
		}

		if config.AdminName == claims.Flag {
			ok, msg := CheckPermission(claims.UserName, c.Request.RequestURI, c.Request.Method)
			if !ok {
				c.JSON(http.StatusOK, response.Error("权限不足:"+msg, 0, nil))
				c.Abort()
			}
		}
		c.Set("username", claims.UserName)

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
	return false, ""
}
