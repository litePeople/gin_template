package router

import (
	"gin_template/controllers"
	"gin_template/modules/routerx"
)

func init() {
	ns := routerx.NewRouterGroup("v1",
		routerx.RouterGroup("api",
			routerx.RouterGroup("enclosure",
				routerx.RouterAdd("upload", &controllers.EnclosureCtl{}, "POST:EnclosureUpload"),
			),
		),
	)
	routerx.RouterInit(ns)
}
