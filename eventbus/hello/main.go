package main

import (
	"eventbus/common"
	"fmt"
)

type Plugin struct{}

func (plugin *Plugin) Initialize(bus common.Bus) {
	bus.Subscribe("hello", plugin.Hello)
	bus.Subscribe("plugins", plugin.Name)
}

func (plugin *Plugin) Hello(common.Event) {
	fmt.Println("Hello, World!")
}

func (plugin *Plugin) Name(common.Event) {
	fmt.Println("Hello")
}
