// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package init

// Injectors from wire.go:

func InitStart(path string) (*connPool, func(), error) {
	v, err := InitConfig(path)
	if err != nil {
		return nil, nil, err
	}
	initServerConfig, cleanup, err := InitMysql(v)
	if err != nil {
		return nil, nil, err
	}
	initConnPool, cleanup2, err := InitServer(initServerConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return initConnPool, func() {
		cleanup2()
		cleanup()
	}, nil
}
