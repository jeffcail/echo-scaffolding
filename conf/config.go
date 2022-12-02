package conf

import (
	"sync"

	"github.com/echo-scaffolding/core/conf/driver"
)

type CoreConfig struct {
	Debug    bool
	HTTPBind string
	Mysql    driver.MysqlConfig
	Redis    driver.RedisConfig
	Logger   struct {
		Path      string
		MaxSize   int
		MaxAge    int
		Compress  bool
		LocalTime bool
	}
}

var Config *CoreConfig

func NewCoreConfig() {
	var once sync.Once
	once.Do(Init)
}

func Init() {
	Config = &CoreConfig{}
}
