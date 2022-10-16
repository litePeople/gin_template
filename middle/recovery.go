package middle

import (
	"gin_template/modules/logx"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
						strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}

			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
			if brokenPipe {
				logx.Errorx("[gin recover]",
					zap.String("path", ctx.Request.URL.Path),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)))

				_ = ctx.Error(err.(error))
				ctx.Abort()
				return
			}

			logx.Errorx("[gin recover]",
				zap.Time("time", time.Now()),
				zap.Any("error", err),
				zap.String("request", string(httpRequest)),
				zap.String("stack", string(debug.Stack())),
			)

			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	ctx.Next()
}
