package utilx

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/fatih/structs"
	"github.com/shopspring/decimal"

	"github.com/jinzhu/now"
	"github.com/pkg/errors"
)

//除法，校验被除数为0时返回接口为0
// return one/two
func ExcepFloat64(one, two float64) float64 {
	if two == 0 {
		return 0
	}
	return one / two
}

//除法，校验被除数为0时返回接口为0
// return one/two
func ExcepInt(one, two int) int {
	if two == 0 {
		return 0
	}
	return one / two
}

//除法，校验被除数为0时返回接口为0
// return one/two
func ExcepInt64(one, two int64) int64 {
	if two == 0 {
		return 0
	}
	return one / two
}

//自定义小数的位数
func Float64Decimal(v float64, decimal int) float64 {
	format := strings.Builder{}
	format.WriteString("%.0")
	format.WriteString(fmt.Sprintf("%d", decimal))
	format.WriteString("f")
	return Float64(fmt.Sprintf(format.String(), v))
}

func MatchArray(id string, ids []string) bool {
	if id == "" {
		return false
	}
	for i := 0; i < len(ids); i++ {
		if id == ids[i] {
			return true
		}
	}
	return false
}
func IntMatchArr(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

const (
	CMReg    = "^1(3[4-9]|4[7]|5[0-27-9]|7[08]|8[2-478])\\d{8}$"
	CUReg    = "^1(3[0-2]|4[5]|5[56]|7[0156]|8[56])\\d{8}$"
	CTReg    = "^1(3[3]|4[9]|53|7[037]|8[019])\\d{8}$"
	MailReg  = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	PhoneReg = "^1[3456789]\\d{9}$"
	PwdReg   = `[a-zA-Z0-9\W_]{8,}`
	IPReg    = `^((\d{1,2}|1[0-9][0-9]|2[0-5][0-5])\.){3}(\d{1,2}|1[0-9][0-9]|2[0-5][0-5])$`
)

//ValidMail 验证邮箱，符合返回true，否则返回false
func ValidMail(mail string) (match bool) {
	match, _ = regexp.MatchString(MailReg, mail)
	return
}

//ValidPhone 验证手机号，符合返回true，否则返回false
func ValidPhone(phone string) (match bool) {
	match, _ = regexp.MatchString(PhoneReg, phone)
	return
}

// ValidPwd 验证高强度密码，，符合返回true，否则返回false
func ValidPwd(pwd string) (match bool) {
	match, _ = regexp.MatchString(PwdReg, pwd)
	return
}

// ValidIP 验证IP，符合返回true，否则返回false
func ValidIP(ip string) (match bool) {
	match, _ = regexp.MatchString(IPReg, ip)
	return
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-08
 * @copyright [tike]
 * @license   [tike]
 * @desc      [获取结构体中字段的名称]
 * @version   [1.0]
 * @param     {[type]}    structName interface{}) ([]string [description]
 */
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-08
 * @copyright [tike]
 * @license   [tike]
 * @desc      [编码转换]
 * @version   [1.0]
 * @param     {[type]}    src        string  [description]
 * @param     {[type]}    srcCode    string  [description]
 * @param     {[type]}    targetCode string) ([]byte       [description]
 */
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-08
 * @copyright [tike]
 * @license   [tike]
 * @desc      [通过结构，反射，拿到对应的字段值，这里是tag的 cloumn值]
 * @version   [1.0]
 * @param     {[type]}    obj  interface{} [description]
 * @param     {[type]}    keys []string)     (map[string]interface{}, error [description]
 */
func StructFindFilter(obj interface{}, keys []string) (map[string]interface{}, error) {
	if !structs.IsStruct(obj) {
		return nil, errors.New("只能转换struct字段")
	}
	fileMap := make(map[string]interface{})
	s := structs.New(obj)
	for _, key := range keys {
		for _, f := range s.Fields() {
			if !f.IsExported() {
				continue
			}
			_, tag := parseStructTag(f.Tag("orm"))
			column, ok := tag["column"]
			if !ok {
				continue
			}
			if column == key {
				if !f.IsZero() {
					fileMap[key] = f.Value()
					break
				}
			}
		}
	}
	return fileMap, nil
}

var supportTag = map[string]int{
	"-":            1,
	"null":         1,
	"index":        1,
	"unique":       1,
	"pk":           1,
	"auto":         1,
	"auto_now":     1,
	"auto_now_add": 1,
	"size":         2,
	"column":       2,
	"default":      2,
	"rel":          2,
	"reverse":      2,
	"rel_table":    2,
	"rel_through":  2,
	"digits":       2,
	"decimals":     2,
	"on_delete":    2,
	"type":         2,
}

// parse struct tag string
func parseStructTag(data string) (attrs map[string]bool, tags map[string]string) {
	attrs = make(map[string]bool)
	tags = make(map[string]string)
	for _, v := range strings.Split(data, ";") {
		if v == "" {
			continue
		}
		v = strings.TrimSpace(v)
		if t := strings.ToLower(v); supportTag[t] == 1 {
			attrs[t] = true
		} else if i := strings.Index(v, "("); i > 0 && strings.Index(v, ")") == len(v)-1 {
			name := t[:i]
			if supportTag[name] == 2 {
				v = v[i+1 : len(v)-1]
				tags[name] = v
			}
		}
	}
	return
}

// parse struct tag string
func GetOrmColumn(data string) (attrs map[string]bool, tags map[string]string) {
	attrs = make(map[string]bool)
	tags = make(map[string]string)
	for _, v := range strings.Split(data, ";") {
		if v == "" {
			continue
		}
		v = strings.TrimSpace(v)
		if t := strings.ToLower(v); supportTag[t] == 1 {
			attrs[t] = true
		} else if i := strings.Index(v, "("); i > 0 && strings.Index(v, ")") == len(v)-1 {
			name := t[:i]
			if supportTag[name] == 2 {
				v = v[i+1 : len(v)-1]
				tags[name] = v
			}
		}
	}
	return
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-11
 * @copyright [tike]
 * @license   [tike]
 * @desc      [判断一个值是否为0值]
 * @version   [1.0]
 * @param     {[type]}    value reflect.Value) (bool [description]
 */
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-08
 * @copyright [tike]
 * @license   [tike]
 * @desc      [如果是int型的 给剪辑掉后面部分]
 * @version   [1.0]
 * @param     {[type]}    s string)       (string [description]
 * @return    {[type]}      [description]
 */
func SplitFloatStr(s string) string {
	value := strings.Split(s, ".")
	if len(value) > 1 {
		return value[0]
	}
	return s
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-08
 * @copyright [tike]
 * @license   [tike]
 * @desc      [把结构 转map]
 * @version   [1.0]
 * @param     {[type]}    obj interface{}) (map[string]interface{}, error [description]
 */
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	fileMap := make(map[string]interface{})
	if !structs.IsStruct(obj) {
		return nil, errors.New("只能转换struct字段")
	}
	s := structs.New(obj)
	for _, f := range s.Fields() {
		column := f.Name()
		if f.IsExported() {
			if !f.IsZero() {
				fileMap[column] = f.Value()
			}
		}
	}
	return fileMap, nil
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-08-14
 * @copyright [tike]
 * @license   [tike]
 * @desc      [取出map的 keys值]
 * @version   [1.0]
 * @param     {[type]}    data map[string]interface{})( []string [description]
 */
func MapGetFileds(data map[string]interface{}) []string {
	fileds := make([]string, 0, 10)
	if data == nil {
		return fileds
	}
	for k, _ := range data {
		fileds = append(fileds, k)
	}
	return fileds
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-10
 * @copyright [tike]
 * @license   [tike]
 * @desc      [判断一个结构，必须字段是否赋值]
 * @version   [1.0]
 * @param     {[type]}    obj        interface{} [description]
 * @param     {[type]}    mustFields []string)     (bool         [description]
 */
func CheckStructMustFiledIsZero(obj interface{}, mustFields []string) bool {
	s := structs.New(obj)
	isZero := true
	for _, f := range s.Fields() {
		if !f.IsExported() {
			continue
		}
		for _, filed := range mustFields {
			if filed == f.Name() {
				if f.IsZero() {
					isZero = false
					return isZero
				}
			}
		}
	}
	return isZero
}

/**
 * @Author    wuxiandashu
 * @DateTime  2018-07-10
 * @copyright [tike]
 * @license   [tike]
 * @desc      [判断一个值是否为 nil 返回空字符串]
 * @version   [1.0]
 * @param     {[type]}    value interface{}) (string [description]
 */
func CheckValueIsNil(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}

// 合并map
func MergeStrKeyMap(mapParam ...map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for _, param := range mapParam {
		for k, v := range param {
			res[k] = v
		}
	}
	return res
}

// TrimUnNumber 去除所有的非数字的字符
func TrimUnNumber(val string) string {
	return regexp.MustCompile("[\u4e00-\u9fa5]").ReplaceAllString(val, "")
}

// FileBase64ToBys base64的文件转换为[]byte
func FileBase64ToBys(fileBase64 string) ([]byte, error) {
	if strings.Contains(fileBase64, ",") { // 如果包含前缀的则把前缀去除
		fileBase64 = fileBase64[strings.Index(fileBase64, ",")+1:]
	}
	return base64.StdEncoding.DecodeString(fileBase64)
}

// FileToMd5 文件的md5
func FileToMd5(fileBys []byte) string {
	fileMd5hash := md5.New()
	fileMd5hash.Write(fileBys)
	return hex.EncodeToString(fileMd5hash.Sum(nil))
}

// RandInt 生成指定区间随机数
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	resRand, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max-min)))
	if err != nil { // crypto 包生成失败时使用match包得随机
		return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min) + min
	}
	return int(resRand.Int64()) + min
}

// 加密密码
func EncryptionPwd(pwd string) string {
	for i := 0; i < 5; i++ {
		pwd = Md5(pwd + "loansys")
	}
	return pwd
}

// FormatStartDate 格式化开始日期
func FormatStartDate(v int64, isM ...bool) int64 {
	if len(isM) > 0 && isM[0] {
		v = v / 1000
	}
	return now.New(time.Unix(v, 0)).BeginningOfDay().Unix()
}

// FormatEndDate 格式化结束日期
func FormatEndDate(v int64, isM ...bool) int64 {
	if len(isM) > 0 && isM[0] {
		v = v / 1000
	}
	return now.New(time.Unix(v, 0)).EndOfDay().Unix()
}

// 浮点数四舍五入
// in 待四舍五入的浮点数;p,精度位数;out,四舍五入后的浮点数
func Round(in float64, p int) (out float64) {
	res, _ := decimal.NewFromFloat(in).Round(int32(p)).Float64()
	return res
}

// 浮点数小数向上取整
// in 待向上取整的浮点数；p为精度位数；out为向上取整后的浮点数
func Ceil(in float64, p int) float64 {
	temp := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(p)))
	res, _ := decimal.NewFromFloat(in).Mul(temp).Ceil().Div(temp).Float64()
	return res
}

// 浮点数小数向下抹零
// in 待向上取整的浮点数；p为精度位数；out为向上取整后的浮点数
func Floor(in float64, p int) float64 {
	temp := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(p)))
	res, _ := decimal.NewFromFloat(in).Mul(temp).Floor().Div(temp).Float64()
	return res
}

//// Round float64保留小数位
//func Round(f float64, n int) float64 {
//	pow10_n := math.Pow10(n)
//	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
//}

// FormatAmount 格式化金额，从元开始，每三位添加一个逗号
func FormatAmount(amount float64) string {
	amountStr := fmt.Sprintf("%+v", amount)
	length := len(amountStr)
	if length < 4 {
		return amountStr
	}
	arr := strings.Split(amountStr, ".") //用小数点符号分割字符串,为数组接收
	length = len(arr[0])
	if length < 4 {
		return amountStr
	}
	count := (length - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length-(i+1)*3] + "," + arr[0][length-(i+1)*3:]
	}

	return strings.Join(arr, ".")
}
