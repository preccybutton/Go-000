// +build wireinject
// The build tag makes sure the stub is not built in the final build.


package init

import (
	"github.com/google/wire"
)

func InitStart(path string) (*connPool, func(), error){
	wire.Build(InitConfig, InitMysql, InitServer)
	return nil, nil, nil
}
