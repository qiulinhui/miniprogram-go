package middleware

import (
	"bookstore/config"
	"bookstore/ext"
	"bookstore/model"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      config.GetString("jwt.realm"),       // 令牌作用域
		Key:        []byte(config.GetString("jwt.key")), // 令牌签发密钥
		Timeout:    time.Hour,                           // 令牌过期时间
		MaxRefresh: time.Hour * (24*7 - 1),              // 令牌有效期刷新时间，令牌刷新的最大有效期是TokenTime + MaxRefresh
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
			c.Set("user", user)
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
			id := int(claims["id"].(float64))
			token := model.RDB.Get(model.Ctx, "token_"+strconv.Itoa(id)).Val()
			if jwt.GetToken(c) != token {
				return nil
			}
			user := new(model.User)
			if model.DB.First(&user, claims["id"]).Error != nil {
				return nil
			}
			return user
		},
		LoginResponse: func(c *gin.Context, code int, message string, times time.Time) {
			user, exists := c.Get("user")
			if exists {
				err := model.RDB.Set(model.Ctx, "token_"+strconv.Itoa(int(user.(*model.User).ID)), message, time.Hour*(24*7-1)).Err()
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": "Redis服务异常，请稍后重试。",
					})
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "服务器异常。",
				})
			}
			c.JSON(code, gin.H{
				"code":   code,
				"token":  message,
				"expire": times,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, message string, times time.Time) {
			claims := jwt.ExtractClaims(c)
			id := int(claims["id"].(float64))
			// 覆盖之前的令牌，保证同一时间只有一个令牌失效
			err := model.RDB.Set(model.Ctx, "token_"+strconv.Itoa(id), message, time.Hour*(24*7-1)).Err()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "Redis服务异常，请稍后重试。",
				})
			}
			c.JSON(code, gin.H{
				"code":   code,
				"token":  message,
				"expire": times,
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			claims := jwt.ExtractClaims(c)
			id := int(claims["id"].(float64))
			model.RDB.Del(model.Ctx, "token_"+strconv.Itoa(id))
			c.JSON(code, gin.H{
				"code": code,
				"msg":  "退出登录",
			})
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
