package service

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/orm-packaging/gorm/database"
)

var (
	NotIn = "not"
	In    = "in"
	Like  = "like"
	Equal = "="
)

type Concrete struct {
}

// Insert ...
func (*Concrete) Insert(i interface{}) (err error) {
	o := database.Connection()
	err = o.Save(i).Error
	return
}

// Read ...
func (*Concrete) Read(out interface{}) (err error) {
	o := database.Connection()
	err = o.Find(out, out).Error
	return
}

// Update ...
func (*Concrete) Update(i interface{}) (err error) {
	o := database.Connection()
	err = o.Model(i).Updates(i).Error
	return
}

// UpdatesByMap update record by map
func (c *Concrete) UpdatesByMap(i interface{}, value, cond map[string]interface{}) (err error) {
	o := database.Connection()
	err = c.queryBuilder(o, cond).Model(i).Update(value).Error
	return
}

// ReadByMap read by cond
func (*Concrete) ReadByMap(out interface{}, cond map[string]interface{}) (err error) {
	o := database.Connection()
	err = o.Find(out, cond).Error
	return
}

// Delete ... this can't constraint bool、int = 0
func (*Concrete) Delete(i interface{}) (err error) {
	o := database.Connection()
	err = o.Unscoped().Where(i).Delete(i).Error
	return
}

// DeleteByMap ...
func (c *Concrete) DeleteByMap(i interface{}, cond map[string]interface{}) (err error) {
	o := database.Connection().Unscoped().Debug()
	err = c.queryBuilder(o, cond).Delete(i).Error
	return
}

// Query ...
func (*Concrete) Query(out, cond interface{}) (err error) {
	o := database.Connection()
	err = o.Find(out, cond).Error
	return
}

// QueryByMap ...
func (c *Concrete) QueryByMap(out interface{}, cond map[string]interface{}) (err error) {
	o := database.Connection()
	err = c.queryBuilder(o, cond).Find(out).Error
	return
}

// queryBuilder generate db query base condition map
func (*Concrete) queryBuilder(db *gorm.DB, cond map[string]interface{}) *gorm.DB {
	if cond == nil {
		return db
	}

	for key, value := range cond {
		if c := condAnalyze(key); len(c) > 1 {
			switch c[1] {
			case NotIn:
				db = db.Not(value)
			case In:
				db = db.Where(fmt.Sprintf("%s %s (?)", c[0], c[1]), value)
			case Like:
				db = db.Where(fmt.Sprintf("%s %s ?", c[0], c[1]), fmt.Sprintf("%%%s%%", value))
			default:
				db = db.Where(fmt.Sprintf("%s %s ?", c[0], c[1]), value)
			}
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	return db
}

func condAnalyze(key string) []string {
	return strings.Split(key, "__")
}
