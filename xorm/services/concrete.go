package services

import (
	"reflect"

	"github.com/go-xorm/builder"
	"github.com/orm-packaging/xorm/database"
)

type ConcreteService struct {
}

func (base ConcreteService) Read(v interface{}) (exist bool, err error) {
	exist, err = database.OrmEngine.Get(v)
	return
}

// insert single or multi  record
func (base ConcreteService) Insert(v interface{}) (err error) {
	_, err = database.OrmEngine.Insert(v)
	return
}

func (base ConcreteService) Update(v, condition interface{}) (err error) {
	_, err = database.OrmEngine.Update(v, condition)
	return
}

// update with bool 、int attribute
func (base ConcreteService) UpdateWithMap(m map[string]interface{}, condition, bean interface{}) (err error) {
	ty := reflect.TypeOf(condition)
	switch ty.Kind() {
	case reflect.Map:
		_, err = database.OrmEngine.Table(bean).Where(builder.Eq(condition.(map[string]interface{}))).Update(m)
	default:
		_, err = database.OrmEngine.Table(bean).Update(m, condition)
	}
	return
}

// condition is map, the where condition of  bool、int attribute will be effect
func (base ConcreteService) Delete(bean, condition interface{}) (err error) {
	ty := reflect.TypeOf(condition)
	switch ty.Kind() {
	case reflect.Map:
		_, err = database.OrmEngine.Table(bean).Where(builder.Eq(condition.(map[string]interface{}))).Unscoped().Delete(bean)
	default:
		_, err = database.OrmEngine.Table(bean).Unscoped().Delete(condition)
	}

	return
}
