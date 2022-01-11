package middleware

import (
	"bookstore/config"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AuthzMiddleware() gin.HandlerFunc {
	adapter, err := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetString("db.username"),
		config.GetString("db.password"),
		config.GetString("db.hostname"),
		config.GetString("db.port"),
		config.GetString("db.database")), true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	m, err := casbinModel.NewModelFromString(`
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
