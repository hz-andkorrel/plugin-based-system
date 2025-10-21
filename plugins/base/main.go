package main

import (
    "plugin"

	"plugins/base/domain"
)

func CreateDataStore() []domain.BaseStudent {
	return []domain.BaseStudent{
		{Name: "Alice", Number: "S001"},
		{Name: "Bob", Number: "S002"},
	}
}

func main() {
    helloPlugin, _ := plugin.Open("./plugins/hello.so")
    sayHello, _ := helloPlugin.Lookup("SayHello")
    sayHello.(func())()
}
