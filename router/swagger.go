package router

import (
	"gin_template/middle"
	"gin_template/modules/ginx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func swagger(engine *ginx.Engine) {
	swag := engine.Group("/swagger")
	swag.Use(middle.Swagger)
	swag.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
