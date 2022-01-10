package middleware

import (
	"bookstore/config"
	"bookstore/ext"
	"bookstore/model"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      config.GetString("jwt.realm"),       // 令牌作用域
		Key:        []byte(config.GetString("jwt.key")), // 令牌签发密钥
		Timeout:    time.Hour,                           // 令牌过期时间
		MaxRefresh: time.Hour,                           // 令牌有效期刷新时间，令牌刷新的最大有效期是TokenTime + MaxRefresh
		// 回调函数，它应该根据登录信息对用户进行身份验证。
		// 必须返回用户数据作为用户标识符，它将存储在Claim Array中。 必需的。
		// 检查错误(e)以确定适当的错误消息。
		Authenticator: func(c *gin.Context) (interface{}, error) {
			type Params struct {
				Code string `json:"code"`
			}
			params := Params{}
			user := new(model.User)
			c.BindJSON(&params)
			auth := ext.WeChat().GetMiniProgram(ext.MiniprogramCfg).GetAuth()
			result, err := auth.Code2Session(params.Code)
			if err != nil {
				return nil, err
			}
			err = user.FindUserByOpenid(result.OpenID)
			if err != nil {
				user.Openid = result.OpenID
				user.SessionKey = result.SessionKey
				err = user.Create()
				if err != nil {
					return nil, err
				}
			}
			err = user.UpdateSessionKey(result.SessionKey)
			if err != nil {
				return nil, err
			}
			return user, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			return data != nil
		},

		// 登录时调用，向webtoken添加其他有效数据
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{"id": v.ID}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "msg": message})
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := new(model.User)
			if model.DB.First(&user, claims["id"]).Error != nil {
				return nil
			}
			return user
		},
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization, query:token ",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + err.Error())
	}
	return authMiddleware
}
