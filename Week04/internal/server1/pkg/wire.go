// +build wireinject
// The build tag makes sure the stub is not built in the final build.


package pkg

import (
	"Go-000/Week04/api"
	"Go-000/Week04/internal/server1/biz"
	"Go-000/Week04/internal/server1/data"
	"Go-000/Week04/internal/server1/service"
	"github.com/google/wire"
)

func InitService() api.GetTestDataServiceServer{
	wire.Build(data.NewTestDataRepo, biz.NewTestDataUsercase, service.NewGetTestDataService)
	return nil
}

//func InitService() *service.GetTestDataService{
//	wire.Build(data.NewTestDataRepo, biz.NewTestDataUsercase, service.NewGetTestDataService)
//	return nil
//}

