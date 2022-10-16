package router

import (
	"gin_template/modules/ginx"
	"gin_template/modules/routerx"
	"os"
	"path/filepath"
)

func Setup(engine *ginx.Engine) {
	dir, _ := os.Executable()
	path := filepath.Dir(dir)
	// 注册swagger
	swagger(engine)
	engine.Static("/static", path+"/static")

	routerx.RouterSetup(engine)
}
