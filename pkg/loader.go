package pkg

import (
	"strconv"
	"sync"

	"github.com/astaxie/beego"
)

var (
	SingleLoader *Loader
	once         sync.Once
)

type Loader struct {
}

func init() {
	once.Do(func() {
		if SingleLoader == nil {
			SingleLoader = new(Loader)
		}
	})
}

// Get key format is key or runmode::key
func (l *Loader) Get(key string, def ...string) (v string) {
	v = beego.AppConfig.String(key)
	if len(v) == 0 && len(def) > 0 {
		v = def[0]
		return
	}

	return
}

func (l *Loader) GetInt64(key string, def ...int64) (v int64) {
	strV := beego.AppConfig.String(key)
	if len(strV) == 0 && len(def) > 0 {
		v = def[0]
		return
	}
	v, _ = strconv.ParseInt(strV, 10, 64)
	return
}
