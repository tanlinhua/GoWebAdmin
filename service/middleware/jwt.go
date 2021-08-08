package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tanlinhua/go-web-admin/pkg/response"
	"github.com/tanlinhua/go-web-admin/service/config"
)

var JwtKey = []byte(config.JwtKey)

type Claims struct {
	UserName string
	Id       int
	jwt.StandardClaims
}

// 生成token
func GetJWT(username string, id int) (bool, string) {
	expireTime := time.Now().Add(1 * time.Hour).Unix() // expireTime = 3600
	claims := &Claims{
		UserName: username,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,        //过期时间
			IssuedAt:  time.Now().Unix(), //生成时间
			Issuer:    "127.0.0.1",       //签名颁发者
			Subject:   "apiToken",        //签名主题
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return false, err.Error()
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

		if err != nil {
			response.New(c).Error(-2, err.Error())
			c.Abort()
			return
		}
		if !token.Valid {
			response.New(c).Error(-2, "TOKEN ERROR")
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

/*
session
熟悉session运行机制的同学都知道，用户的session数据以file或缓存（redis、memcached）等方式存储在服务器端，客户端浏览器cookie中只保存sessionid。
服务器端session属于集中存储，数量不大的情况下，没什么问题，当用户数据逐渐增多到一程度，就会给服务端管理和维护带来大的负担。
session有两个弊端：
1、无法实现跨域。
2、由于session数据属于集中管理里，量大的时候服务器性能是个问题。
优点：
1、session存在服务端，数据相对比较安全。
2、session集中管理也有好处，就是用户登录、注销服务端可控。

cookie
cookie也是一种解决网站用户认证的实现方式，用户登录时，服务器会发送包含登录凭据的Cookie到用户浏览器客户端，浏览器会将Cookie的key/value保存用户本地（内存或硬盘），用户再访问网站，浏览器会发送cookie信息到服务器端，服务器端接收cookie并解析来维护用户的登录状态。
cookie避免session集中管理的问题，但也存在弊端：
1、跨域问题。
2、数据存储在浏览器端，数据容易被窃取及被csrf攻击，安全性差。
优点：
1、相对于session简单，不用服务端维护用户认证信息。
2、数据持久性。

jwt
jwt通过json传输，php、java、golang等很多语言支持，通用性比较好，不存在跨域问题。传输数据通过数据签名相对比较安全。
客户端与服务端通过jwt交互，服务端通过解密token信息，来实现用户认证。不需要服务端集中维护token信息，便于扩展。当然jwt也有其缺点。
缺点：
1、用户无法主动登出，只要token在有效期内就有效。这里可以考虑redis设置同token有效期一直的黑名单解决此问题。
2、token过了有效期，无法续签问题。可以考虑通过判断旧的token什么时候到期，过期的时候刷新token续签接口产生新token代替旧token。

jwt设置有效期
可以设置有效期，加入有效期是为了增加安全性，即token被黑客截获，也只能攻击较短时间。设置有效期就会面临token续签问题，解决方案如下

通常服务端设置两个token
Access Token：添加到 HTTP 请求的 header 中，进行用户认证，请求接口资源。
refresh token：用于当 Access Token过期后，客户端传递refresh token刷新 Access Token续期接口，获取新的Access Token和refresh token。
其有效期比 Access Token有效期长。


小结
服务端生成的jwt返回客户端可以存到cookie也可以存到localStorage中（相比cookie容量大），存在cookie中需加上 HttpOnly 的标记，可以防止 XSS) 攻击。
尽量用https带证书网址访问。
session和jwt没有绝对好与不好，各有其擅长的应用环境，请根据实际情况选择。
*/
