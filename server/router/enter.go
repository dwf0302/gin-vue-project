package router

import (
	"go-vue-project/router/example"
	"go-vue-project/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
