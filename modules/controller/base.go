package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Ctx *gin.Context
}

// 在请求之前
// 返回fasle 终止请求
func (c *Controller) OnRequest() bool {

	return true
}
