package centrifugox

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// GenTokenByExpire 根据过期时间生成token
// accountId 账户id
// s  过期时间，单位为秒
// token token
// err 错误信息
func (c *client) GenTokenByExpire(accountId int64, s int64) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       strconv.FormatInt(accountId, 10),
		"iat":       time.Now().Unix(),                                             // 发出令牌的时间戳
		"expire_at": time.Now().Add(time.Second * time.Duration(s)).Unix(),         // 令牌在6个小时后过期
		"exp":       time.Now().Add(time.Second*time.Duration(s) + 24*3600).Unix(), // 连接在一天后过期
	})
	token, err = jwtToken.SignedString(c.tokenHmacSecretKeyBys)
	return
}

// GenToken 生成token
// accountId 账户id
// token token
// err 错误信息
func (c *client) GenToken(accountId int64) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       strconv.FormatInt(accountId, 10),
		"iat":       time.Now().Unix(),                     // 发出令牌的时间戳
		"expire_at": time.Now().Add(time.Hour * 6).Unix(),  // 令牌在6个小时后过期
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 连接在一天后过期
	})
	token, err = jwtToken.SignedString(c.tokenHmacSecretKeyBys)
	return
}

// GenPrivateToken 生成私有频道的token
// clientId 允许该clientId的连接订阅私人频道，客户端在连接上centrifugo的时候会得到clientId。
// channel 频道名，注意：私有频道名格式：$[频道名]
// token 生成的token
// err 错误信息
func (c *client) GenPrivateToken(clientId, channel string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client":    clientId,
		"channel":   channel,
		"iat":       time.Now().Unix(),                     // 发出令牌的时间戳
		"expire_at": time.Now().Add(time.Hour * 6).Unix(),  // 令牌在6个小时后过期
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 连接在一天后过期
	})
	token, err = jwtToken.SignedString(c.tokenHmacSecretKeyBys)
	return
}

// Publish 推送消息
// channelName 通道名，使用 constant.CENTRIFUGO__CHANNEL__xxx
// value 需要推送的信息
// error 错误信息
func (c *client) Publish(channelName string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		content []byte
		err     error
	)

	content, err = json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "序列化失败")
	}

	err = ctfg.cli.Publish(ctx, channelName, content)
	if err != nil {
		return errors.Wrap(err, "推送失败")
	}
	return nil
}

// GenTokenByExpire 根据过期时间生成token
// accountId 账户id
// s  过期时间，单位为秒
// token token
// err 错误信息
func GenTokenByExpire(accountId int64, s int64) (token string, err error) {
	return ctfg.GenTokenByExpire(accountId, s)
}

// GenToken 生成token
// accountId 账户id
// token token
// err 错误信息
func GenToken(accountId int64) (token string, err error) {
	return ctfg.GenToken(accountId)
}

// GenPrivateToken 生成私有频道的token
// clientId 允许该clientId的连接订阅私人频道，客户端在连接上centrifugo的时候会得到clientId。
// channel 频道名，注意：私有频道名格式：$[频道名]
// token 生成的token
// err 错误信息
func GenPrivateToken(clientId, channel string) (token string, err error) {
	return ctfg.GenPrivateToken(clientId, channel)
}

// Publish 推送消息
// channelName 通道名，使用 constant.CENTRIFUGO__CHANNEL__xxx
// value 需要推送的信息
// error 错误信息
func Publish(channelName string, value interface{}) error {
	return ctfg.Publish(channelName, value)
}
