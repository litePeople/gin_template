package gormx

import (
	"context"
	"database/sql"
	"fmt"
	"gin_template/modules/component"
	"gin_template/modules/config"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var cli client

func init() {
	component.RegComponent(&cli)
}

const (
	ConfSectionMysql = "mysql"
)

type client struct {
	gdb *gorm.DB `ini:"-"`
	// 地址
	Host string
	// 端口
	Port uint
	// 名称
	Name string
	// 用户
	User string
	// 密码
	Password string
	// 表前缀
	TablePrefix string
	// 闲置的
	MaxIdle uint
	// 活跃的
	MaxActive uint
	// 超时的，单位为秒
	IdleTimeout uint
	// 活跃的存活时间，单位为秒
	ActiveLifetime uint
}

func (c *client) Setup(cfg *ini.File) (err error) {
	var (
		connStr  string
		db       *sql.DB
		dialetor gorm.Dialector
		config   *gorm.Config
	)

	err = cfg.Section(ConfSectionMysql).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析mysql模块的配置失败")
		return
	}

	connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		c.User, c.Password, c.Host, c.Port, c.Name)

	db, err = sql.Open("mysql", connStr)
	if err != nil {
		err = errors.Wrap(err, "[gormx] 打开mysql失败")
		return
	}

	err = ping(db)
	if err != nil {
		err = errors.Wrap(err, "[gormx] ping失败")
		return
	}
	db.SetConnMaxIdleTime(time.Duration(c.IdleTimeout) * time.Second)
	db.SetMaxOpenConns(int(c.MaxActive))
	db.SetMaxIdleConns(int(c.MaxIdle))
	db.SetConnMaxLifetime(time.Duration(c.ActiveLifetime) * time.Second)

	//配置Dialetor
	dialetor = mysql.New(mysql.Config{Conn: db})
	//配置config
	config = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键迁移
	}

	config.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  getLogLevel(), // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)

	//创建db
	c.gdb, err = gorm.Open(dialetor, config)
	if err != nil {
		err = errors.Wrap(err, "[gorm] 打开gorm失败")
		return
	}
	return
}

// ping ping数据库，测试数据库是否真正可以连接
func ping(db *sql.DB) (err error) {
	dbPingCtx, cancelF := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelF()

	done := make(chan struct{})
	go func() {
		err = db.PingContext(dbPingCtx)
		done <- struct{}{}
	}()
	select {
	case <-dbPingCtx.Done():
		err = dbPingCtx.Err()
	case <-done:
	}
	return err
}

func getLogLevel() logger.LogLevel {
	switch config.GetAPP().LogLevel {
	case "debug":
		return logger.Info
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "fatal", "panic":
		return logger.Silent
	default:
		return logger.Info
	}
}
