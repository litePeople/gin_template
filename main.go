package main

import (
	"context"
	"fmt"
	"gin_template/consts"
	_ "gin_template/docs/swagger"
	"gin_template/middle"
	_ "gin_template/models"
	_ "gin_template/modules/alioss"
	_ "gin_template/modules/appletsx"
	"gin_template/modules/component"
	"gin_template/modules/config"
	"gin_template/modules/ginx"
	_ "gin_template/modules/gormx"
	"gin_template/modules/logx"
	_ "gin_template/modules/redisx"
	_ "gin_template/modules/validatorx"
	"gin_template/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	time.Local = consts.CSTSH
}

// @title Swagger  API
// @version 1.0
// @description 小程序的的接口文档
// @description 错误码的含义：
// @description 错误码：1001，含义：参数错误
// @description 错误码：1002，含义：业务层错误
// @description 错误码：1003，含义：未登录
// @description 错误码：1004，含义：wxid重复绑定
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https
// @host www.xxx.cn
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	err := component.SetupComponents()
	if nil != err {
		panic(err)
	}

	gin.SetMode(config.GetServer().RunMode)

	engine := ginx.New()
	// 设置nginx透传过来的IP
	engine.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP", "X-real-ip", "RemoteIP"}

	// 将IP设置到context中
	// engine.Use(middle.IP2Ctx)

	// 确保发生错误时也不崩溃
	engine.Use(middle.Recovery)
	// 请求日志打印
	engine.Use(middle.Log)
	// 请求头
	engine.Use(middle.Header)

	// 初始化路由
	router.Setup(engine)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetServer().HttpPort),
		Handler:        engine,
		ReadTimeout:    time.Duration(config.GetServer().ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.GetServer().WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 10, // 10M
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGKILL)

	select {
	case <-quit:
		logx.Info("Shutdown Server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logx.Fatal("Server Shutdown:", err)
		}

		logx.Infox("Server existing!!!")
	}
}
