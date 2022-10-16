package logx

import (
	"gin_template/modules/component"
	"gin_template/modules/config"
	"os"

	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cli client

func init() {
	component.RegComponent(&cli)
}

type client struct {
	levelCtl zap.AtomicLevel
	logger   *logxLogger
}

// Setup 必须要在conf的setup之后
func (c *client) Setup(config *ini.File) (err error) {
	cli.levelCtl = zap.NewAtomicLevelAt(zap.DebugLevel)
	c.changeLevel()
	stdEncoderCfg := zap.NewDevelopmentEncoderConfig()
	stdEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder        // 采用 2006-01-02T15:04:05.000Z0700 时间格式
	stdEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // 采用彩色显示不同level的日志
	stdEncoder := zapcore.NewConsoleEncoder(stdEncoderCfg)
	stdCore := zapcore.NewCore(stdEncoder, os.Stdout, c.levelCtl)

	c.logger = &logxLogger{
		logger: zap.New(stdCore, zap.AddCaller(),
			zap.AddCallerSkip(2),
			zap.AddStacktrace(zapcore.ErrorLevel)),
	}
	return
}

// PriorityOn 模块的加载优先级
func (c *client) PriorityOn() int {
	return 999998
}

func (c *client) changeLevel() {
	switch config.GetAPP().LogLevel {
	case "debug":
		c.levelCtl.SetLevel(zap.DebugLevel)
	case "info":
		c.levelCtl.SetLevel(zap.InfoLevel)
	case "error":
		c.levelCtl.SetLevel(zap.ErrorLevel)
	case "warn":
		c.levelCtl.SetLevel(zap.WarnLevel)
	case "panic":
		c.levelCtl.SetLevel(zap.PanicLevel)
	case "fatal":
		c.levelCtl.SetLevel(zap.FatalLevel)
	}
}
