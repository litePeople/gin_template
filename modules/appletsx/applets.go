package appletsx

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/subscribeMessage"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/support"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/auth"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/phoneNumber"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/security"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment"
	"github.com/pkg/errors"
)

// PlainData 用户信息/手机号信息
type PlainData struct {
	OpenID      string `json:"openId"`
	UnionID     string `json:"unionId"`
	NickName    string `json:"nickName"`
	Gender      int    `json:"gender"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	AvatarURL   string `json:"avatarUrl"`
	Language    string `json:"language"`
	OpenGID     string `json:"openGId"`
	CountryCode string `json:"countryCode"`
	Watermark   struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

func GetAuth() *auth.Client {
	return applets.miniProgram.Auth
}

func GetPhoneNumber() *phoneNumber.Client {
	return applets.miniProgram.PhoneNumber
}

func GetSecurity() *security.Client {
	return applets.miniProgram.Security
}

func GetEncryptor() *miniProgram.Encryptor {
	return applets.miniProgram.Encryptor
}

func GetSubscribe() *subscribeMessage.Client {
	return applets.miniProgram.SubscribeMessage
}

func GetMsgCallbackToken() string {
	return applets.MsgCallbackToken
}

func GetPay() *payment.Payment {
	return applets.payment
}

func GetMchID() string {
	return applets.MchID
}

func GetAppID() string {
	return applets.AppID
}

// DecryptUserInfo 解密用户信息
func DecryptUserInfo(encrypted string, sessionKey string, iv string) (dta *PlainData, err error) {
	var (
		bys      []byte
		cryptErr *support.CryptError
	)

	bys, cryptErr = GetEncryptor().DecryptData(encrypted, sessionKey, iv)
	if nil != cryptErr {
		err = errors.Errorf("code:%d,msg:%s", cryptErr.ErrCode, cryptErr.ErrMsg)
		return
	}

	dta = &PlainData{}

	err = json.Unmarshal(bys, dta)
	if nil != err {
		return
	}

	return
}
