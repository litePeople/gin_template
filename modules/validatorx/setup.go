package validatorx

import (
	"gin_template/modules/component"

	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var cli client

func init() {
	cli.vali = validator.New()
	component.RegComponent(&cli)
}

type client struct {
	vali *validator.Validate
}

func (c *client) Setup(config *ini.File) (err error) {
	err = c.vali.RegisterValidation("notemptystring", verifyEmptyString)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 notemptystring 校验失败")
		return
	}
	err = c.vali.RegisterValidation("orderby", verifyOrderBy)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 orderby 校验失败")
		return
	}
	err = c.vali.RegisterValidation("oneof", verifyOneOf)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 oneof 校验失败")
		return
	}
	err = c.vali.RegisterValidation("lt", verifyLt)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 lt 校验失败")
		return
	}
	err = c.vali.RegisterValidation("lte", verifyLte)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 lte 校验失败")
		return
	}
	err = c.vali.RegisterValidation("gt", verifyGt)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 gt 校验失败")
		return
	}
	err = c.vali.RegisterValidation("gte", verifyGte)
	if nil != err {
		err = errors.Wrap(err, "[validatorx] 初始化 gte 校验失败")
		return
	}
	return
}
