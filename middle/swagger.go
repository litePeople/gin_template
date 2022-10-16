package middle

import (
	"crypto/subtle"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Swagger(ctx *gin.Context) {
	var (
		creds = map[string]string{
			"admin": "admin@kblife",
		}
		auth        string
		isEqualFold bool
		lower       = func(b byte) byte {
			if 'A' <= b && b <= 'Z' {
				return b + ('a' - 'A')
			}
			return b
		}
		authContentBys       []byte
		authContent          string
		authContentFirstByte int
		ok                   bool
		err                  error
	)
	const prefix = "Basic "

	defer func() {
		if ok {
			ctx.Next() // 继续往下执行
			return
		}
		ctx.Header("WWW-Authenticate", `Basic realm="Access to api"`)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}()

	auth = ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return
	}

	if len(auth) < len(prefix) {
		return
	}

	if len(auth[:len(prefix)]) != len(prefix) {
		return
	}

	isEqualFold = true
	for i := 0; i < len(auth[:len(prefix)]); i++ {
		if lower(auth[:len(prefix)][i]) != lower(prefix[i]) {
			isEqualFold = false
			break
		}
	}
	if !isEqualFold {
		return
	}

	authContentBys, err = base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	authContent = string(authContentBys)
	authContentFirstByte = strings.IndexByte(authContent, ':')
	if authContentFirstByte < 0 {
		return
	}
	// user authContent[:authContentFirstByte]
	// pass authContent[authContentFirstByte+1:]

	credPass, credUserOk := creds[authContent[:authContentFirstByte]]
	if !credUserOk || subtle.ConstantTimeCompare([]byte(authContent[authContentFirstByte+1:]), []byte(credPass)) != 1 {
		return
	}

	ok = true
	return
}
