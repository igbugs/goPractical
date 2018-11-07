package datasource

import (
	"github.com/go-xorm/xorm"
	"sync"
	"superstar/conf"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Password, c.Host, c.Port, c.DBName)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("datasource.InstanceMaster create connection failed, err:%v", err)
		return nil
	}

	engine.ShowSQL(false)
	engine.SetTZLocation(conf.SysTimeLocation)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Password, c.Host, c.Port, c.DBName)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("datasource.InstanceSlave create connection failed, err:%v", err)
		return nil
	}

	engine.ShowSQL(false)
	engine.SetTZLocation(conf.SysTimeLocation)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	slaveEngine = engine
	return slaveEngine
}