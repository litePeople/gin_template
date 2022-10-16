package gormx

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LockForUpdate 数据库锁
// clause.Expression
func LockForUpdate() clause.Expression {
	return clause.Locking{Strength: "UPDATE"}
}

// Paginator 查询分页器 使用GORM特性实现的针对GORM的查询分页器
// nthPage 下一页，从1开始
// pageSize 每页大小
// func(db *gorm.DB) *gorm.DB
func Paginator(nthPage, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nthPage == 0 {
			nthPage = 1
		}
		offset := (nthPage - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// PreloadUser 预加载用户，只有最基础的信息：id、nickname、avatar、gender
func PreloadUser() func(gdb *gorm.DB) *gorm.DB {
	return func(gdb *gorm.DB) *gorm.DB {
		return gdb.Select("id", "nickname", "avatar", "gender")
	}
}

func AutoMigrate(dst ...interface{}) error {
	return cli.gdb.AutoMigrate(dst...)
}

// GetDB 获取gorm的db
func GetDB() *gorm.DB {
	return cli.gdb
}

// GetDBCtx 获取gorm的db
func GetDBCtx(ctx context.Context) *gorm.DB {
	if result, ok := ctx.Value(gormCTXKey).(*gorm.DB); !ok {
		return cli.gdb
	} else {
		return result
	}
}

func Table(name string, args ...interface{}) (tx *gorm.DB) {
	return cli.gdb.Table(name, args...)
}

func Select(query string, args ...interface{}) (tx *gorm.DB) {
	return cli.gdb.Select(query, args...)
}

// 通过map对象进行传参数
func Condit(condit map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condit["is_deleted = ?"] = 0
		for key, value := range condit {
			db = db.Where(key, value)
		}
		return db
	}
}

func TransactionCtx(ctx context.Context, fc func(txCtx context.Context) error, opts ...*sql.TxOptions) (err error) {
	return GetDBCtx(ctx).Transaction(func(tx *gorm.DB) error {
		return fc(SetDBToCtx(ctx, tx))
	}, opts...)
}

func Create(value interface{}) (tx *gorm.DB) {
	return cli.gdb.Create(value)
}

func CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
	return cli.gdb.CreateInBatches(value, batchSize)
}

func TablePrefix() string {
	return cli.TablePrefix
}

const gormCTXKey = "gormKey"

// SetDBToCtx 设置db到context
func SetDBToCtx(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, gormCTXKey, db)
}
