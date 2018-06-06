package database

import (
	"fmt"
	"sync"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/orm-packaging/pkg"
)

var (
	OrmEngine *xorm.Engine
	once      sync.Once
)

var (
	dataSource       = "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	DEFAULT_USER     = "root"
	DEFAULT_PASSWORD = "root"
	DEFAULT_HOST     = "127.0.0.1"
	DEFAULT_PORT     = "3306"
	DEFAULT_DATABASE = "xxx"
	MAX_IDLE_CONN    = 30
	MAX_OPEN_CONN    = 50
)

func init() {
	NewOrmEngine()
}

type DatabaseInfo struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
}

func (d *DatabaseInfo) verbose() {
	d.User = pkg.GetVariable("DB_USER", DEFAULT_USER)
	d.Password = pkg.GetVariable("DB_PASSWORD", DEFAULT_PASSWORD)
	d.Host = pkg.GetVariable("DB_URL", DEFAULT_HOST)
	d.Port = pkg.GetVariable("DB_PORT", DEFAULT_PORT)
	d.Database = pkg.GetVariable("MYSQL_DATABASE", DEFAULT_DATABASE)
}

func (d *DatabaseInfo) Builder() string {
	d.verbose()
	return fmt.Sprintf(dataSource, d.User, d.Password, d.Host, d.Port, d.Database)
}

func NewDatabaseInfo() *DatabaseInfo {
	return new(DatabaseInfo)
}

func NewOrmEngine() {
	var err error
	once.Do(func() {
		OrmEngine, err = xorm.NewEngine("mysql", NewDatabaseInfo().Builder())
		if err = OrmEngine.Ping(); err != nil {
			panic(fmt.Sprintf("init orm engine failed, error info is %s", err.Error()))
		}
		setOrmEngineInfo()
	})
	return
}

func setOrmEngineInfo() {
	OrmEngine.SetMaxIdleConns(MAX_IDLE_CONN)
	OrmEngine.SetMaxOpenConns(MAX_OPEN_CONN)
	OrmEngine.SetMapper(core.GonicMapper{})
	OrmEngine.ShowSQL(false)
}
