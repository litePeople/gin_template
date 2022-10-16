package utilx

import (
	"fmt"
	"strings"
)

var sqlSpecialcharacters = []string{
	"%", "_",
}

// SqlEscape 转义特殊字符，防止sql注入特殊符号导致可以直接获取所有列表
func SqlEscape(val string) string {
	for _, item := range sqlSpecialcharacters {
		val = strings.ReplaceAll(val, item, fmt.Sprintf("\\%s", item))
	}
	return val
}
