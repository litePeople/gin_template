package routerx

// LinkNamespace used as link action
type LinkNamespace func(*Namespace)

// Namespace is store all the info
type Namespace struct {
	prefix   string
	Child    []*Namespace
	handlers *ControllerRegister
}

func (ns *Namespace) Namespace(s *Namespace) {
	ns.Child = append(ns.Child, s)
	// ns.handlers.Add(s.handlers.routers...)
}

// ControllerRegister containers registered router rules, controller handlers and filters.
type ControllerRegister struct {
	routers []*Routerx
}

func (c *ControllerRegister) Add(rs ...*Routerx) {
	c.routers = append(c.routers, rs...)
}

type Routerx struct {
	FullPath string      // 全地址 最终地址
	Path     string      // 路由地址
	Target   interface{} // 处理控制器
	Method   string      // 处理方法 GET POST PUT
	Handler  string      // 处理的方法名称
}

// NewRouterGroup get new Namespace
func NewRouterGroup(prefix string, params ...LinkNamespace) *Namespace {
	ns := &Namespace{
		prefix:   prefix,
		Child:    make([]*Namespace, 0, 20),
		handlers: NewControllerRegister(),
	}
	for _, p := range params {
		p(ns)
	}
	return ns
}

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister{
		routers: make([]*Routerx, 0, 10),
	}
}
