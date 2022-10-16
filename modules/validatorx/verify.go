package validatorx

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"strings"
)

// verifyEmptyString 校验空字符串
// fl 需要校验的值的相关信息
// bool true-空字符串，false-不是空字符串
func verifyEmptyString(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

// verifyOrderBy 校验排序字段
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyOrderBy(fl validator.FieldLevel) bool {
	orderBy := fl.Field().String()

	if "" == orderBy { // 排序字段为空则不校验格式
		return true
	}
	// 切割排序的参数
	orderByNameMp := make(map[string]struct{}, 10)
	for _, item := range strings.Split(fl.Param(), " ") {
		orderByNameMp[item] = struct{}{}
	}
	for _, filed := range strings.Split(orderBy, ",") { // 以逗号切割每个字段的排序
		if "" == filed {
			return false
		}

		filedInfo := strings.Split(filed, " ")
		if len(filedInfo) > 2 {
			return false
		}

		if len(filedInfo) == 2 {
			if strings.ToLower(filedInfo[1]) != "desc" && strings.ToLower(filedInfo[1]) != "asc" {
				return false
			}
		}

		if _, ok := orderByNameMp[filedInfo[0]]; !ok {
			return false
		}

	}

	return true
}

// verifyOneOf 校验候选值
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyOneOf(fl validator.FieldLevel) bool {
	// 切割指定指的参数,并组装为map
	oneofMp := make(map[string]struct{}, 10)
	for _, item := range strings.Split(fl.Param(), " ") {
		oneofMp[item] = struct{}{}
	}
	values := make([]string, 0, 10)

	switch fl.Field().Interface().(type) {
	case []int, []int8, []int16, []int32, []int64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, fmt.Sprintf("%d", fl.Field().Index(idx).Int()))
		}
	case []uint, []uint8, []uint16, []uint32, []uint64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, fmt.Sprintf("%d", fl.Field().Index(idx).Uint()))
		}
	case []string:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			if "" != fl.Field().String() {
				values = append(values, fl.Field().Index(idx).String())
			}
		}
	case int, int8, int16, int32, int64:
		values = append(values, fmt.Sprintf("%d", fl.Field().Int()))
	case uint, uint8, uint16, uint32, uint64:
		values = append(values, fmt.Sprintf("%d", fl.Field().Uint()))
	case float32, float64:
		values = append(values, decimal.NewFromFloat(fl.Field().Float()).String())
	case string:
		if "" != fl.Field().String() {
			values = append(values, fl.Field().String())
		}
	default:
		return false
	}

	for _, item := range values {
		if _, ok := oneofMp[item]; !ok {
			return false
		}
	}

	return true
}

// verifyLt 校验小于
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyLt(fl validator.FieldLevel) bool {
	lt, err := decimal.NewFromString(fl.Param())
	if nil != err {
		return false
	}

	values := make([]decimal.Decimal, 0, 10)
	switch fl.Field().Interface().(type) {
	case []int, []int8, []int16, []int32, []int64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(fl.Field().Index(idx).Int()))
		}
	case []uint, []uint8, []uint16, []uint32, []uint64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(int64(fl.Field().Index(idx).Uint())))
		}
	case int, int8, int16, int32, int64:
		values = append(values, decimal.NewFromInt(fl.Field().Int()))
	case uint, uint8, uint16, uint32, uint64:
		values = append(values, decimal.NewFromInt(int64(fl.Field().Uint())))
	case []float32, []float64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromFloat(fl.Field().Index(idx).Float()))
		}
	case float32, float64:
		values = append(values, decimal.NewFromFloat(fl.Field().Float()))
	default:
		return false
	}

	for _, item := range values {
		sign := lt.Sub(item).Sign()
		if sign == -1 || sign == 0 {
			return false
		}
	}
	return true
}

// verifyLte 校验小于等于
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyLte(fl validator.FieldLevel) bool {
	lt, err := decimal.NewFromString(fl.Param())
	if nil != err {
		return false
	}

	values := make([]decimal.Decimal, 0, 10)
	switch fl.Field().Interface().(type) {
	case []int, []int8, []int16, []int32, []int64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(fl.Field().Index(idx).Int()))
		}
	case []uint, []uint8, []uint16, []uint32, []uint64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(int64(fl.Field().Index(idx).Uint())))
		}
	case int, int8, int16, int32, int64:
		values = append(values, decimal.NewFromInt(fl.Field().Int()))
	case uint, uint8, uint16, uint32, uint64:
		values = append(values, decimal.NewFromInt(int64(fl.Field().Uint())))
	case []float32, []float64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromFloat(fl.Field().Index(idx).Float()))
		}
	case float32, float64:
		values = append(values, decimal.NewFromFloat(fl.Field().Float()))
	default:
		return false
	}

	for _, item := range values {
		sign := lt.Sub(item).Sign()
		if sign == -1 {
			return false
		}
	}
	return true
}

// verifyGt 校验大于
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyGt(fl validator.FieldLevel) bool {
	lt, err := decimal.NewFromString(fl.Param())
	if nil != err {
		return false
	}

	values := make([]decimal.Decimal, 0, 10)
	switch fl.Field().Interface().(type) {
	case []int, []int8, []int16, []int32, []int64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(fl.Field().Index(idx).Int()))
		}
	case []uint, []uint8, []uint16, []uint32, []uint64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(int64(fl.Field().Index(idx).Uint())))
		}
	case int, int8, int16, int32, int64:
		values = append(values, decimal.NewFromInt(fl.Field().Int()))
	case uint, uint8, uint16, uint32, uint64:
		values = append(values, decimal.NewFromInt(int64(fl.Field().Uint())))
	case []float32, []float64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromFloat(fl.Field().Index(idx).Float()))
		}
	case float32, float64:
		values = append(values, decimal.NewFromFloat(fl.Field().Float()))
	default:
		return false
	}

	for _, item := range values {
		sign := item.Sub(lt).Sign()
		if sign == -1 || sign == 0 {
			return false
		}
	}
	return true
}

// verifyGte 校验大于等于
// fl 需要校验的值的相关信息
// bool true-符合，false-不符合
func verifyGte(fl validator.FieldLevel) bool {
	lt, err := decimal.NewFromString(fl.Param())
	if nil != err {
		return false
	}

	values := make([]decimal.Decimal, 0, 10)
	switch fl.Field().Interface().(type) {
	case []int, []int8, []int16, []int32, []int64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(fl.Field().Index(idx).Int()))
		}
	case []uint, []uint8, []uint16, []uint32, []uint64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromInt(int64(fl.Field().Index(idx).Uint())))
		}
	case int, int8, int16, int32, int64:
		values = append(values, decimal.NewFromInt(fl.Field().Int()))
	case uint, uint8, uint16, uint32, uint64:
		values = append(values, decimal.NewFromInt(int64(fl.Field().Uint())))
	case []float32, []float64:
		for idx := 0; idx < fl.Field().Len(); idx++ {
			values = append(values, decimal.NewFromFloat(fl.Field().Index(idx).Float()))
		}
	case float32, float64:
		values = append(values, decimal.NewFromFloat(fl.Field().Float()))
	default:
		return false
	}

	for _, item := range values {
		sign := item.Sub(lt).Sign()
		if sign == -1 {
			return false
		}
	}
	return true
}
