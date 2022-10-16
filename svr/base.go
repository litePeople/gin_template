package svr

import "context"

const (
	loginAccIdCTXKey = "loginAccIdKey"
	loginAccIPCTXKey = "loginAccIPKey"
)

//GetLoginAccIdCtx 一般把用户登录ID设置在ctx里面
func GetLoginAccIdCtx(ctx context.Context) int64 {
	if result, ok := ctx.Value(loginAccIdCTXKey).(int64); !ok {
		return 0
	} else {
		return result
	}
}
func GetLoginAccIPCtx(ctx context.Context) string {
	if result, ok := ctx.Value(loginAccIPCTXKey).(string); !ok {
		return ""
	} else {
		return result
	}
}
