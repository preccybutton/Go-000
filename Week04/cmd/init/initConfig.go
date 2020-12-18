package init

import "errors"

type mysqlConfig struct {
	maxIdle		int
	maxConn		int
	maxLifetime	int
	addr string
}

type serverConfig struct {
	lis string
}

type initConfig struct {
	sqlc map[string]*mysqlConfig
	serc *serverConfig
}

var cfg *initConfig


func (c *initConfig) Sqlc() map[string]*mysqlConfig{
	if c != nil{
		return c.sqlc
	}
	return nil
}




func (c *initConfig) SetSerc(data *serverConfig) {
	c.serc = data
}



func (c *initConfig) Serc() *serverConfig{
	if c != nil{
		return c.serc
	}
	return nil
}

func InitConfig(path string)(map[string]*mysqlConfig, error  ){
	if path == "aaa"{
		cfg = &initConfig{
			sqlc: map[string]*mysqlConfig{"test": {maxConn: 10, maxIdle: 5, maxLifetime: 12, addr:"test:123456@tcp(172.19.4.225:3306)/mysql?charset=utf8"}},
			serc:&serverConfig{lis:"1041.141.41.1:8000"}}
		return cfg.Sqlc(), nil
	}
	return nil, errors.New("aaa")

}
