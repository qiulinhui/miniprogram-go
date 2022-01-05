package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
)

func LoadAuthorizationRules() (rules grbac.Rules, err error) {
	return rules, nil
}

func QueryRolesByHeaders(header http.Header) (roles []string, err error) {
	// 在这里实现你的逻辑
	// ...
	// 这个逻辑可能是从请求的Headers中获取token，并且根据token从数据库中查询用户的相应角色。
	return roles, err
}

func Authorization() gin.HandlerFunc {
	rbac, err := grbac.New(grbac.WithLoader(LoadAuthorizationRules, time.Minute))
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		roles, err := QueryRolesByHeaders(c.Request.Header)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		state, _ := rbac.IsRequestGranted(c.Request, roles)
		if !state.IsGranted() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
