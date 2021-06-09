package svc

import (
	"book/service/user/api/internal/config"
	"book/service/user/api/internal/middleware"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"book/service/user/model"
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn,c.CacheRedis),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
