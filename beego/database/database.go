package database

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/orm-packaging/pkg"
)

var (
	DefaultUser     = "root"
	DefaultPwd      = "root"
	DefaultHost     = "localhost"
	DefaultPort     = "3306"
	DefaultDatabase = "xxx"
)

func init() {
	runmode := pkg.SingleLoader.Get("runmode")
	username := pkg.SingleLoader.Get(fmt.Sprintf("%s::%s", runmode, "username"), DefaultUser)
	password := pkg.SingleLoader.Get(fmt.Sprintf("%s::%s", runmode, "password"), DefaultPwd)
	host := pkg.SingleLoader.Get(fmt.Sprintf("%s::%s", runmode, "host"), DefaultHost)
	port := pkg.SingleLoader.Get(fmt.Sprintf("%s::%s", runmode, "port"), DefaultPort)
	database := pkg.SingleLoader.Get(fmt.Sprintf("%s::%s", runmode, "database"), DefaultDatabase)

	orm.RegisterDataBase("default", "mysql", dataSource(username, password, host, port, database), 30)
	RegistryModel()
}

func dataSource(username, password, host, port, database string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?utf-8", username, password, host, port, database)
	fmt.Println(dsn)
	return dsn
}

func RegistryModel() {
	orm.RegisterModel()

	orm.RunSyncdb("default", false, false)
}
