package middlewares

import (
	"app/app"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AuthzMiddleware() gin.HandlerFunc {
	adapter, err := gormadapter.NewAdapterByDB(app.DB)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	[policy_definition]
	p = sub, obj, act
	[policy_effect]
	e = some(where(p.eft == allow))
	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`)
	if err != nil {
		log.Fatalf("error:model:%s", err)
	}

	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("error: enforcer:%s", err)
	}
	return authz.NewAuthorizer(e)
}
