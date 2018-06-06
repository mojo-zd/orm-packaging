package database

import (
	"errors"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/orm-packaging/pkg"
)

var (
	dbClient        *gorm.DB
	DefaultPort     = "3306"
	DefaultAddress  = "127.0.0.1"
	DefaultUser     = "root"
	DefaultPassword = "root123"
)

func init() {
	database := &Database{
		User:     pkg.GetVariable("DB_USER", DefaultUser),
		Password: pkg.GetVariable("DB_PASSWORD", DefaultPassword),
		Port:     pkg.GetVariable("DB_PORT", DefaultPort),
		Address:  pkg.GetVariable("DB_URL", DefaultAddress),
		Database: pkg.GetVariable("DATABASE", ""),
	}
	if err := OpenConnection(database); err != nil {
		panic(fmt.Sprintf("数据库连接失败, %s", err.Error()))
	}
	logrus.Println("数据库连接成功")
}

type Database struct {
	User     string
	Password string
	Charset  string
	Database string
	Address  string
	Port     string
}

func OpenConnection(database *Database) error {
	connection, err := initDatabase(database)
	if err != nil {
		return err
	}

	dbClient, err = gorm.Open("mysql", connection)
	logrus.Println(connection)
	return err
}

func initDatabase(database *Database) (connection string, err error) {
	if database == nil {
		err = errors.New("未设置数据库信息")
		return
	}
	connection = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

	if len(database.User) == 0 {
		database.User = "root"
	}
	if len(database.Password) == 0 {
		database.Password = "root"
	}

	if len(database.Database) == 0 {
		err = errors.New("未指定数据库,创建数据库连接失败")
		return
	}
	connection = fmt.Sprintf(connection, database.User, database.Password, database.Address+":"+database.Port, database.Database)
	return
}

func Connection() *gorm.DB {
	if dbClient == nil {
		panic("dbclient is nil")
	}

	return dbClient
}
