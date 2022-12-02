// Code generated by CodeGenerator. Not generate if exist
//
// Source: account.proto
// Time: 2022-11-30 15:50:38

package svc

import "ProjectX/service/account/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	// TODO: 连接数据库，创建协程，设置缓冲等
	return &ServiceContext{
		Config: c,
	}
}