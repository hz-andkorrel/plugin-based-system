package main

import (
    "plugin"

    "plugins/base/common"
	"plugins/base/domain"
)

func CreateDataStore() []common.Student {
	return []common.Student{
        domain.NewBaseStudent("Alice", "S001"),
        domain.NewBaseStudent("Bob", "S002"),
	}
}

func main() {
    students := CreateDataStore()

    helloPlugin, _ := plugin.Open("./plugins/hello.so")
    sayHello, _ := helloPlugin.Lookup("SayHello")
    sayHello.(func(common.Student))(students[0])
}
