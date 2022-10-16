package validatorx

import (
	"github.com/pkg/errors"
	"testing"
)

func init() {
	cli.Setup(nil)
}
func TestTagNotEmptyString(t *testing.T) {
	_ = cli.Setup(nil)
	type tag struct {
		Name string `json:"name" validate:"required,max=50,notemptystring"`
	}

	allEmpty := tag{}
	err := Struct(allEmpty)
	if nil != err {
		t.Log(errors.Wrap(err, "所有为空"))
	}

	emptyString := tag{
		Name: " ",
	}
	err = Struct(emptyString)
	if nil != err {
		t.Log(errors.Wrap(err, "空字符串"))
	}

	normal := tag{
		Name: "姓名",
	}
	err = Struct(normal)
	if nil != err {
		t.Log(errors.Wrap(err, "正确"))
	}
}

func TestTagOrderBy(t *testing.T) {
	type tag struct {
		OrderBy string `form:"orderBy" validate:"orderby=price updateTime receiveNum" format:"字段名+排序方式" example:"price desc,updateTime,receiveNum desc"`
	}

	emptyTag := tag{
		OrderBy: "",
	}

	err := Struct(emptyTag)
	if nil != err {
		t.Log(errors.Wrap(err, "空字符串"))
	}

	failTag := tag{
		OrderBy: "createTime desc",
	}
	err = Struct(failTag)
	if nil != err {
		t.Log(errors.Wrap(err, "错误的排序方式"))
	}

	normal := tag{
		OrderBy: "price desc,updateTime,receiveNum desc",
	}

	err = Struct(normal)
	if nil != err {
		t.Log(errors.Wrap(err, "正确的排序方式"))
	}

}

func TestTagOneOf(t *testing.T) {
	var (
		state1 = struct {
			State int `form:"state" validate:"oneof=0 10"`
		}{
			State: 10,
		}
		state2 = struct {
			State int64 `form:"state" validate:"oneof=0 10"`
		}{
			State: 10000,
		}
		state3 = struct {
			State []int `form:"state" validate:"oneof=0 10"`
		}{
			State: []int{10},
		}
		state5 = struct {
			State []string `form:"state" validate:"oneof=0 10"`
		}{
			State: []string{"10"},
		}
		state6 = struct {
			State string `form:"state" validate:"oneof=0 10"`
		}{
			State: "10",
		}
	)

	err := Struct(state1)
	if nil != err {
		t.Log(err)
	}
	err = Struct(state2)
	if nil != err {
		t.Log(err)
	}
	err = Struct(state3)
	if nil != err {
		t.Log(err)
	}
	err = Struct(state5)
	if nil != err {
		t.Log(err)
	}
	err = Struct(state6)
	if nil != err {
		t.Log(err)
	}
	errs := Map(map[string]interface{}{
		"gender": 0.003053214000,
	}, map[string]interface{}{
		"gender": "oneof=0 1 2",
	})
	for _, item := range errs {
		t.Log(item)
	}
}

func TestTagLt(t *testing.T) {
	type tag struct {
		Width float64 `form:"width" validate:"lt=10.865"`
	}

	failTag := tag{
		Width: 10.87,
	}

	err := Struct(failTag)
	if nil != err {
		t.Log(err)
	}

	normal := tag{
		Width: 10,
	}
	err = Struct(normal)
	if nil != err {
		t.Log(err)
	}

}

func TestTagArrLt(t *testing.T) {
	type tag struct {
		Width []float64 `form:"width" validate:"lt=10.865"`
	}

	failTag := tag{
		Width: []float64{10.87},
	}

	err := Struct(failTag)
	if nil != err {
		t.Log(err)
	}

	normal := tag{
		Width: []float64{10},
	}
	err = Struct(normal)
	if nil != err {
		t.Log(err)
	}

}
