package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Header(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, token")
	ctx.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	ctx.Header("Content-Type", "application/json")
	reqMethod := ctx.Request.Method
	ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	ctx.Header("Connection", "close")
	ctx.Header("X-Frame-Options", "DENY")
	ctx.Header("X-Content-Type-Options", "nosniff")
	ctx.Header("X-XSS-Protection", "1; mode=block")
	if ctx.Request.TLS != nil {
		ctx.Header("Strict-Transport-Security", "max-age=31536000")
	}
	if reqMethod == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}
	ctx.Next()
}
