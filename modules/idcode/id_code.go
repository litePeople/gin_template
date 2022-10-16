package idcode

import (
	"math/rand"
	"strings"
	"time"
)

var (
	base    = "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ" // 进制的包含字符, string类型
	decimal = uint64(32)                         // 进制长度
	pad     = "F"                                // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
	minLen  = 6                                  // code最小长度
)

// Id2Code id转code
func Id2Code(id uint64) string {
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % decimal
		id = id / decimal
		res += string(base[mod])
	}
	resLen := len(res)
	if resLen < minLen {
		res += pad
		for i := 0; i < minLen-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(base[rand.Intn(int(decimal))])
		}
	}
	return res
}

// Code2Id code转id
func Code2Id(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)

	//var baseArr [] byte = []byte(base)
	baseArr := []byte(base)       // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}

	// 查找补位字符的位置
	isPad := strings.Index(code, pad)
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == pad {
			continue
		}
		index := baseRev[code[i]]
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= decimal
		}
		// pow 类型为 float64 , 类型转换太麻烦, 所以自己循环实现pow的功能
		//res += float64(index) * math.Pow(float64(32), float64(2))
		res += uint64(index) * b
		r++
	}
	return res
}
