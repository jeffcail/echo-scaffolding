package boot

import "github.com/echo-scaffolding/conf"

//InitIni
func InitIni() {
	conf.ParseIniConfig()
}
