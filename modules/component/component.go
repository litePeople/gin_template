package component

import (
	"github.com/go-ini/ini"
)

type Component interface {
	// Setup 安装
	Setup(config *ini.File) (err error)
}

type Priority interface {
	// PriorityOn 优先级
	PriorityOn() int
}

type componentItem struct {
	// 模块
	component Component
	// 优先级
	priority int
}

type componentItems []componentItem

func (pi componentItems) Len() int           { return len(pi) }
func (pi componentItems) Less(i, j int) bool { return pi[i].priority > pi[j].priority }
func (pi componentItems) Swap(i, j int)      { pi[i], pi[j] = pi[j], pi[i] }
