package routerx

import (
	"fmt"
	"gin_template/modules/ginx"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var ROUTERS []*Routerx

func init() {
	ROUTERS = make([]*Routerx, 0, 100)
}

// 路由分组
func RouterGroup(prefix string, params ...LinkNamespace) LinkNamespace {
	return func(ns *Namespace) {
		n := NewRouterGroup(prefix, params...)
		ns.Namespace(n)
	}
}

// 新增路由
// path 路径
// object 控制器对象
// method 处理方法 eg: GET,POST:Login
func RouterAdd(path string, object interface{}, method string) LinkNamespace {
	return func(ns *Namespace) {
		if path == "" {
			path = "/"
		}
		methods := strings.Split(method, ":")
		reqMethod := "*"
		handler := ""
		if len(method) <= 0 {
			panic(fmt.Sprintf("路径:%s,方法不能为空", path))
		}
		if len(methods) == 1 {
			handler = methods[0]
		}
		if len(methods) == 2 {
			reqMethod = methods[0] //请求方法
			handler = methods[1]   //处理方法名称
		}
		routerx := &Routerx{
			Path:    path,
			Target:  object,
			Method:  reqMethod,
			Handler: handler,
		}
		// logs.Notice("加入路径:%s", path)
		ns.handlers.Add(routerx)
	}
}

func RouterInit(groups ...*Namespace) {
	for _, group := range groups {
		RouterParse(group, []string{})
	}
}

// 路由解析
func RouterParse(group *Namespace, paths []string) {
	//子级编辑
	paths = append(paths, strings.TrimLeft(group.prefix, "/"))
	if len(group.Child) > 0 {
		// logs.Notice("路径相加:%+v,当前:%+v,group:%+v", paths, group.prefix, group)
		for _, child := range group.Child {
			RouterParse(child, paths)
		}
	}

	// logs.Notice("整合分组中的路径数组:%+v,child:%+v,路由数组:%d", paths, group.Child, len(group.handlers.routers))
	for _, router := range group.handlers.routers {
		rootPaths := append(paths, strings.TrimLeft(router.Path, "/"))
		path := strings.Join(rootPaths, "/")
		// logs.Notice("最终路径:%s,router:%+v", path, router)
		router.FullPath = path
		// RouterHandle(router)
		ROUTERS = append(ROUTERS, router) //添加到路由地址
	}
}

// Context 扩展gin的context
type Context struct {
	*gin.Context
}

// RouuterHandler 处理请求函数
func RouterHandler(router *Routerx) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0, 2)
	handler := func(ctx *gin.Context) {
		// logs.Notice("处理函数 路径:%s,处理对象:%s,路由结构:%+v", router.FullPath, router.Method, router)
		// ctx.JSON(200, "收到请求。等待处理")
		routerHandlerTarget(ctx, router)
	}
	handlers = append(handlers, handler)
	return handlers
}

func routerHandlerTarget(ctx *gin.Context, router *Routerx) {
	targetType := reflect.TypeOf(router.Target)
	targetValue := reflect.ValueOf(router.Target)
	// 需要是指针类型
	if targetType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("路径:%s,目标对象不是指针类型", router.FullPath))
	}
	//判断是不是nil
	if targetValue.IsNil() {
		panic(fmt.Sprintf("路径:%s,目标对象为空指针", router.FullPath))
	}

	//判断一下原始类型是否是结构体
	originType := targetValue.Elem().Type()
	// logs.Notice("处理函数值原始类型:%+v,对象名称:%s", originType.Kind(), originType.Name())
	if originType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("路径:%s,目标对象必须是结构体", router.FullPath))
	}
	methodNums := targetType.NumMethod()

	elemType := targetType.Elem()
	// elemVal := targetValue.Elem()

	// logs.Notice("处理函数值方法数量:%d,字段数量:%d", methodNums, targetType.Elem().NumField())
	methodHash := make(map[string]bool)
	for i := 0; i < methodNums; i++ {
		method := targetType.Method(i)
		methodHash[method.Name] = true
		// if method.Name == router.Handler {
		// 	rValues := make([]reflect.Value, 0)
		// 	targetValue.MethodByName(router.Handler).Call(rValues)
		// }
		// logs.Notice("处理函数值类型:方法名称:%s,路由绑定方法:%s,值:%+v", method.Name, router.Handler, method.Type)
	}

	//检测绑定的方法是否存在
	if _, ok := methodHash[router.Handler]; !ok {
		panic(fmt.Sprintf("路由：%s,绑定对象名称:%s,方法:%s,不存在", router.FullPath, originType.Name(), router.Handler))
	}

	ctxValue := reflect.ValueOf(ctx)

	//需要new一个新对象 出来
	elemVal := reflect.New(targetType.Elem())
	reflectTargetField(elemType, elemVal.Elem(), ctxValue)

	//最后 调用新对象的相应的处理方法
	rValues := make([]reflect.Value, 0)

	//调用请求处理之前方法
	refValues := elemVal.MethodByName("OnRequest").Call(rValues)
	if len(refValues) > 0 {
		if refValues[0].Kind() != reflect.Bool || !refValues[0].Bool() {
			ctx.Abort()
			return
		}
	}
	elemVal.MethodByName(router.Handler).Call(rValues)
}

// 深层次遍历迭代 是对象的结构体
func reflectTargetField(elemType reflect.Type, elemVal reflect.Value, ctxValue reflect.Value) {
	// logs.Notice("值:%+v", elemVal)
	num := elemType.NumField()
	for i := 0; i < num; i++ {
		field := elemType.Field(i)
		fieldName := field.Name
		fieldValue := elemVal.FieldByIndex(field.Index)
		// logs.Notice("处理函数值类型:结构名称:%s,字段名称:%s,值:%+v,类型:%+v", elemType.Name(), fieldName, fieldValue, field.Type.Kind())
		if fieldName == "Ctx" && fieldValue.CanSet() {
			fieldValue.Set(ctxValue)
		}
		if field.Type.Kind() == reflect.Struct {
			// logs.Notice("处理函数值类型:字段:%s,还是对象，需要继续遍历：%+v", fieldName, field.Type.Kind())
			reflectTargetField(field.Type, fieldValue, ctxValue)
		}
	}
}

// 路由初始化
func RouterSetup(engine *ginx.Engine) {
	for _, router := range ROUTERS {
		// logs.Notice("路由初始化：%s", router.FullPath)
		// engine.GET(router.FullPath, RouterHandler(router)...)

		//需要判断注册什么类型的路由

		//任意请求类型
		if router.Method == "*" {
			engine.Any(router.FullPath, RouterHandler(router)...)
			continue
		}
		//支持多个 如果 POST,GET

		methods := strings.Split(router.Method, ",")
		for _, method := range methods {
			if strings.ToUpper(method) == "GET" {
				engine.GET(router.FullPath, RouterHandler(router)...)
			}
			if strings.ToUpper(method) == "POST" {
				engine.POST(router.FullPath, RouterHandler(router)...)
			}
			if strings.ToUpper(method) == "PUT" {
				engine.PUT(router.FullPath, RouterHandler(router)...)
			}
			if strings.ToUpper(method) == "DELETE" {
				engine.DELETE(router.FullPath, RouterHandler(router)...)
			}
			if strings.ToUpper(method) == "OPTIONS" {
				engine.OPTIONS(router.FullPath, RouterHandler(router)...)
			}
			if strings.ToUpper(method) == "HEAD" {
				engine.HEAD(router.FullPath, RouterHandler(router)...)
			}
		}
	}
}
