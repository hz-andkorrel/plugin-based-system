package main

import (
	"eventbus/bus/domain"
	"eventbus/common"
	"fmt"
)

func main() {
	// Create an instance of the eventbus and the plugin-respository
	eventBus := domain.NewBus()
	pluginRegistry := domain.NewPluginRegistry()

	// Register plugins
	pluginRegistry.AutoDiscoverPlugins()
	pluginRegistry.InitializePlugins(eventBus)

	// Joink some events to the bus
	fmt.Println("\nPublishing hello event")
	helloEvent := common.Event{Name: "hello", Data: nil}
	eventBus.Publish(helloEvent)

	fmt.Println("\nPublishing bye event")
	byeEvent := common.Event{Name: "bye", Data: nil}
	eventBus.Publish(byeEvent)

	fmt.Println("\nPublishing plugins event")
	pluginEvent := common.Event{Name: "plugins", Data: nil}
	eventBus.Publish(pluginEvent)
}
