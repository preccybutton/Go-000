package init

import (
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"time"
)


var conn map[string]*xorm.Engine

func InitMysql(sqlc map[string]*mysqlConfig)(*serverConfig, func(), error){
	var err error
	conn = make(map[string]*xorm.Engine)
	fn := func() {
		for _, v := range conn{
			v.Close()
		}
	}
	for k, v := range sqlc{
		err = initSingle(k, " mysql", v)
		if err != nil{
			return nil, fn, err
		}
	}
	return cfg.Serc(), fn, err
}

func initSingle(engineName, dbType string, cfg  *mysqlConfig) error{
	var err error
	if conn[engineName], err = xorm.NewEngine(dbType, cfg.addr); err != nil{
		err = errors.Wrapf(err,"create xorm engine failed , source: %s, err: %v", dbType, cfg.addr)
		return err
	}
	conn[engineName].SetMaxOpenConns(cfg.maxConn)
	conn[engineName].SetMaxIdleConns(cfg.maxIdle)
	conn[engineName].SetConnMaxLifetime(time.Duration(cfg.maxLifetime)*time.Hour)
	if err = conn[engineName].Ping(); err != nil{
		err = errors.Wrapf(err,"ping  db  failed ,  source: %s, err: %v", dbType, cfg.addr)
		return err
	}
	return err
}