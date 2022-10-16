package ginx

import "github.com/gin-gonic/gin"

// Engine 扩展gin的engine
type Engine struct {
	*gin.Engine
}

func New() *Engine {
	return &Engine{
		Engine: gin.New(),
	}
}

func Default() *Engine {
	return &Engine{
		Engine: gin.Default(),
	}
}
