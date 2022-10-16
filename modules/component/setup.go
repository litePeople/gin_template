package component

import (
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
	"sort"
)

type RunEndSetup func() (err error)

// 所有的模块
var components = make([]Component, 0, 20)
var endSetups = make([]RunEndSetup, 0, 20)

// RegComponent 注册组件
func RegComponent(component Component) {
	components = append(components, component)
}

// RegRunEndSetup 注册运行于结束setup的
func RegRunEndSetup(setup RunEndSetup) {
	endSetups = append(endSetups, setup)
}

// SetupComponents 安装所有的组件
func SetupComponents() (err error) {
	var (
		config *ini.File
	)
	config, err = loadConfig()
	if nil != err {
		err = errors.Wrap(err, "加载配置")
		return
	}

	items := make(componentItems, 0, len(components))
	for _, cmp := range components {
		item := componentItem{component: cmp}
		if priority, ok := cmp.(Priority); ok {
			item.priority = priority.PriorityOn()
		}

		items = append(items, item)
	}
	sort.Sort(items)

	for _, item := range components {
		err = item.Setup(config)
		if nil != err {
			err = errors.Wrap(err, "安装组件失败")
			return
		}
	}

	for _, item := range endSetups {
		err = item()
		if nil != err {
			err = errors.Wrap(err, "run end setup失败")
			return
		}
	}

	return
}

// loadConfig 加载配置
func loadConfig() (config *ini.File, err error) {
	var (
		executablePath string
	)
	executablePath, err = os.Executable()
	if nil != err {
		err = errors.Wrap(err, "获取执行路径失败")
		return
	}
	executablePath = path.Dir(executablePath)
	config, err = ini.LooseLoad(filepath.Join(executablePath, "conf/app.ini"),
		filepath.Join(executablePath, "conf/app.local.ini"))
	if err != nil {
		err = errors.Wrap(err, "配置文件不存在")
		return
	}
	return
}
