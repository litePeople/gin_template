package controllers

import (
	"errors"
	"gin_template/modules/controller"
	"gin_template/modules/tokenx"
	"gin_template/modules/validatorx"
	"net/http"
	"reflect"
)

type BaseCtl struct {
	controller.Controller
	UserId int64
}

// Resp 响应信息
type Resp struct {
	// 状态码
	// 200-正确
	Code int `json:"code"`
	// 错误消息
	Message string `json:"message"`
	// 返回数据
	Data interface{} `json:"data,omitempty"`
}

// 在接口请求到来之时，优先执行本方法
func (ctl *BaseCtl) OnRequest() bool {
	ctl.GetUserId()
	return true
}

// JSONS 返回正确的json
func (ctl *BaseCtl) JSONS(obj interface{}) {
	ctl.Ctx.JSON(http.StatusOK, Resp{
		Code: 200,
		Data: obj,
	})
}

// JSONE 返回错误的json
func (ctl *BaseCtl) JSONE(code int, err error) {
	ctl.Ctx.JSON(http.StatusOK, Resp{
		Code:    code,
		Message: err.Error(),
	})
}

// ParseJSON 转换body为json
func (ctl *BaseCtl) ParseJSON(obj interface{}) (err error) {
	err = ctl.Ctx.ShouldBindJSON(obj)
	if nil != err {
		return
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	// 如果是切片或者数组，那么需要一个一个子项校验
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			err = validatorx.Struct(val.Index(i).Interface())
			if nil != err {
				return err
			}
		}
		return nil
	}

	if val.Kind() == reflect.Struct {
		// 校验参数
		return validatorx.Struct(obj)
	}

	return
}

// ParseQuery 转换query参数
func (ctl *BaseCtl) ParseQuery(obj interface{}) (err error) {
	err = ctl.Ctx.ShouldBindQuery(obj)
	if nil != err {
		return
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	// 如果是切片或者数组，那么需要一个一个子项校验
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			err = validatorx.Struct(val.Index(i).Interface())
			if nil != err {
				return err
			}
		}
		return nil
	}
	// 校验参数
	return validatorx.Struct(obj)
}

// ParseForm 转换表单参数
func (ctl *BaseCtl) ParseForm(obj interface{}) (err error) {
	err = ctl.Ctx.Request.ParseForm()
	if nil != err {
		return
	}
	err = ctl.Ctx.Request.ParseMultipartForm(1 << 20)
	if err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return
	}

	err = ctl.Ctx.ShouldBind(obj)
	if nil != err {
		return
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	// 如果是切片或者数组，那么需要一个一个子项校验
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			err = validatorx.Struct(val.Index(i).Interface())
			if nil != err {
				return err
			}
		}
		return nil
	}
	// 校验参数
	return validatorx.Struct(obj)
}

const (
	loginAccIdCTXKey = "loginAccIdKey"
	loginAccIPCTXKey = "loginAccIPKey"
)

// 从token获取用户ID，并且设置到context里面
func (ctl *BaseCtl) GetUserId() int64 {
	token, _ := ctl.Ctx.GetQuery("token")
	if token == "" {
		token = ctl.Ctx.GetHeader("token")
	}
	loginAccId, err := tokenx.GetTokenUserId(token)
	if nil != err {
		return 0
	}
	ctl.Ctx.Set(loginAccIdCTXKey, loginAccId) // 设置登录账号id
	return loginAccId
}

func (ctl *BaseCtl) ParseIPToCtx() {
	ctl.Ctx.Set(loginAccIPCTXKey, ctl.Ctx.ClientIP()) // 设置登录账号的IP
}
