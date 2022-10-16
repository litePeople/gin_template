package models

import (
	"gin_template/modules/component"
	"gin_template/modules/gormx"

	"github.com/pkg/errors"
)

func init() {
	component.RegRunEndSetup(Setup)
}

func Setup() (err error) {
	err = gormx.AutoMigrate()
	if nil != err {
		err = errors.Wrap(err, "自动迁移表失败")
		return
	}
	return
}
