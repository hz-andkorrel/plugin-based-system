package main

import (
	"eventbus/common"
	"fmt"
)

type Plugin struct{}

func NewPlugin() common.Plugin {
	return &Plugin{}
}

func (plugin *Plugin) Initialize(bus common.Bus) {
	bus.Subscribe("hello", func(event common.Event) {
		fmt.Println("Hello from anonymous function!")
	})

	bus.Subscribe("plugins", plugin.Name)
}

func (plugin *Plugin) Hello(common.Event) {
	fmt.Println("Hello, World!")
}

func (plugin *Plugin) Name(common.Event) {
	fmt.Println("Hello")
}
