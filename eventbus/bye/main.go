package main

import (
	"eventbus/common"
	"fmt"
)

type Plugin struct{}

func (plugin *Plugin) Initialize(bus common.Bus) {
	bus.Subscribe("bye", plugin.Bye)
	bus.Subscribe("plugins", plugin.Name)
}

func (plugin *Plugin) Bye(common.Event) {
	fmt.Println("Bye, World!")
}

func (plugin *Plugin) Name(common.Event) {
	fmt.Println("Bye")
}
