package middle

import (
	"gin_template/consts"
	"gin_template/modules/logx"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	ctx.Next()

	end := time.Now()
	latency := end.Sub(start)
	end = end.UTC()

	if len(ctx.Errors) > 0 {
		for _, e := range ctx.Errors.Errors() {
			logx.Errorx("[gin log]", zap.String("error", e))
		}
	} else {
		fields := []zapcore.Field{
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		fields = append(fields, zap.String("time", end.Format(consts.YYYYMMDDHHMMSS)))

		logx.Infox("[gin log]", fields...)
	}
}
