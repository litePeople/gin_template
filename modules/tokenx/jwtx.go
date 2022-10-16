package tokenx

import (
	"fmt"
	"gin_template/modules/utilx"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Generate 生成jwt
func (c *client) Generate(userId int64) (string, error) {
	return c.GenerateWithExp(userId, time.Now().Add(7*24*time.Hour))
}

// GenerateWithExp 生成jwt，并且定制过期时间
func (c *client) GenerateWithExp(userId int64, exp time.Time) (string, error) {
	claimsMap := jwt.MapClaims{
		"exp":          exp.Unix(), // 默认添加过期时间为7天
		tokenUserIdKey: userId,
	}

	// 生成hs 256 得 jwt
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)
	return tokenClaims.SignedString([]byte(c.TokenSigned))
}

// Resolve 解析jwt
func (c *client) Resolve(jwtStr string) (map[string]interface{}, error) {
	res, err := c.parseData(jwtStr)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("jwt格式错误:%s", err.Error())
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("jwt已过期:%s", err.Error())
			} else {
				return nil, fmt.Errorf("jwt错误:%s", err.Error())
			}
		} else {
			return nil, fmt.Errorf("jwt错误:%s", err.Error())
		}
	}
	return res, nil
}

// parseData 解析token数据
func (c *client) parseData(jwtStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.TokenSigned), nil
	})

	if token == nil {
		return nil, err
	}

	res := make(map[string]interface{})
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return res, nil
	}
	for k, v := range token.Claims.(jwt.MapClaims) {
		res[k] = v
	}
	return res, nil
}

// Generate 生成jwt
func Generate(userId int64) (string, error) {
	return cli.Generate(userId)
}

// GenerateWithExp 生成jwt，并且定制过期时间
func GenerateWithExp(userId int64, exp time.Time) (string, error) {
	return cli.GenerateWithExp(userId, exp)
}

// Resolve 解析jwt
func Resolve(jwtStr string) (map[string]interface{}, error) {
	return cli.Resolve(jwtStr)
}

// GetTokenUserId 获取token中的用户id，如果失败则返回0
func GetTokenUserId(token string) (int64, error) {
	dtaMp, err := Resolve(token)
	if nil != err {
		return 0, err
	}

	userId, ok := dtaMp[tokenUserIdKey]
	if !ok {
		return 0, nil
	}
	return utilx.Int64(fmt.Sprintf("%+v", userId)), nil
}

const tokenUserIdKey = "user_id"
