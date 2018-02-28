package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/orm-packaging/pagination"
	"github.com/orm-packaging/tool"
)

type ConcreteService struct {
}

// Read read record by cols
func (base ConcreteService) Read(v interface{}, cols ...string) (err error) {
	err = orm.NewOrm().Read(v, cols...)
	return
}

// Insert insert new record
func (base ConcreteService) Insert(v interface{}) (err error) {
	_, err = orm.NewOrm().Insert(v)
	return
}

// InsertMulti insert multi record
// bulk is the length of v
func (base ConcreteService) InsertMulti(v interface{}) (err error) {
	_, err = orm.NewOrm().InsertMulti(1, v)
	return
}

// Update update v by cols
func (base ConcreteService) Update(v interface{}, cols ...string) (err error) {
	_, err = orm.NewOrm().Update(v, cols...)
	return
}

// UpdateWithMap update single instance with map
// you can get result record from table
func (base ConcreteService) UpdateWithMap(condition, m map[string]interface{}, table interface{}) (err error) {
	query := orm.NewOrm().QueryTable(table)
	_, err = queryBuilder(query, condition).Update(m)
	if err != nil {
		return
	}
	if _, ok := condition["id"]; !ok {
		return
	}
	tool.SetValue(table, "id", condition["id"])
	err = base.Read(table)
	return
}

// ReadByQuery read record by beego orm's query which query is complex
func (base ConcreteService) ReadByQuery(table interface{}, condition map[string]interface{}) (obj interface{}, err error) {
	obj = tool.NewInstanceSetValue(table, "id", condition["id"])
	query := orm.NewOrm().QueryTable(table)
	query = queryBuilder(query, condition)
	err = query.One(&obj)
	return
}

// InsertOrUpdate
func (base ConcreteService) InsertOrUpdate(v interface{}) (err error) {
	_, err = orm.NewOrm().InsertOrUpdate(v)
	return
}

// Delete
func (base ConcreteService) Delete(v interface{}, cols ...string) (err error) {
	_, err = orm.NewOrm().Delete(v, cols...)
	return
}

// Count
func (base ConcreteService) Count(condition map[string]interface{}, table interface{}) (count int64) {
	o := orm.NewOrm()
	query := o.QueryTable(table)

	query = queryBuilder(query, condition)
	count, _ = query.Count()
	return
}

// Exist
func (base ConcreteService) Exist(condition map[string]interface{}, table interface{}) (exist bool) {
	o := orm.NewOrm()
	query := o.QueryTable(table)
	exist = queryBuilder(query, condition).Exist()
	return
}

// QueryElement query record by page
func (base ConcreteService) QueryElement(page *pagination.Pagination, condition map[string]interface{}, table interface{}) orm.QuerySeter {
	o := orm.NewOrm()
	page.SetTotal(base.Count(condition, table))

	query := o.QueryTable(table)
	return queryBuilder(query, condition).Limit(page.PageSize, page.Offset())
}

func queryBuilder(query orm.QuerySeter, condition map[string]interface{}) orm.QuerySeter {
	if condition == nil {
		return query
	}
	for key, value := range condition {
		query = query.Filter(key, value)
	}
	return query
}
