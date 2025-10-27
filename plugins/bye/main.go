package main

import (
	"plugins/common"
)

type Plugin struct{}

func NewPlugin() common.Plugin {
	return &Plugin{}
}

func (plugin *Plugin) Hello() string {
	return "ByePlugin"
}
