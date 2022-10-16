package config

import (
	"gin_template/modules/component"

	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

var cfg config

func init() {
	component.RegComponent(&cfg)
}

const (
	ConfSectionApp    = "app"
	ConfSectionServer = "server"
)

type config struct {
	server Server
	app    APP
}

// Server 服务配置
type Server struct {
	// 运行模式
	// debug-测试模式
	// release-发布模式
	RunMode string
	// http端口
	HttpPort int
	// 服务读超时，单位为秒
	ReadTimeout uint
	// 服务写超时，单位为秒
	WriteTimeout uint
}

type APP struct {
	// 日志级别
	// debug-测试
	// info-信息
	// warn-警告
	// panic-崩溃
	// fatal-退出
	LogLevel string
}

// Setup 安装
func (c *config) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionApp).MapTo(&c.app)
	if nil != err {
		err = errors.Wrap(err, "解析app模块的配置失败")
	}
	if "debug" != c.app.LogLevel &&
		"info" != c.app.LogLevel &&
		"warn" != c.app.LogLevel &&
		"panic" != c.app.LogLevel &&
		"fatal" != c.app.LogLevel {
		err = errors.Wrap(err, "校验app的LogLevel失败")
	}

	err = config.Section(ConfSectionServer).MapTo(&c.server)
	if nil != err {
		err = errors.Wrap(err, "解析server模块的配置失败")
	}

	if "debug" != c.server.RunMode &&
		"release" != c.server.RunMode {
		err = errors.Wrap(err, "校验server的RunMode失败")
	}
	return
}

// PriorityOn 模块的加载优先级
func (c *config) PriorityOn() int {
	return 999999
}

// GetAPP 获取app参数
func (c *config) GetAPP() APP {
	return c.app
}

// GetServer 获取服务参数
func (c *config) GetServer() Server {
	return c.server
}

// GetAPP 获取app参数
func GetAPP() APP {
	return cfg.app
}

// GetServer 获取服务参数
func GetServer() Server {
	return cfg.server
}
