package utilx

import (
	"strconv"
	"strings"
	"unsafe"
)

// Str2bytes 字符串转换字节数组
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str 字节数组转换字符串，注意转换过后不可以直接截取，会panic
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// IsEmpty 判断value是否为空字符串
func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}

// IsNotEmpty 判断value是否为非空字符串
func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func Int64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func Float64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func Bool(str string) bool {
	i, _ := strconv.ParseBool(str)
	return i
}
