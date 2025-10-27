# Research: Creating a plugin based system

This repository holds the demo's for several implementations of plugin systems.
The findings will be discussed in this document.
Each heading describes a single solution.


## Directly implementing go-plugins

Go offers a build-in solution for plug-ins.
By supplying the `-buildmode=plugin` flag, the compiler will create a `.so` file.
This file can be imported in another project using the `plugin` library.
The `main.go` file of a plugin can have the following function in it.

```go
func SayHello() {
    student := domain.ClassStudent{ Name: "Joseph", Number: "S001" }
    fmt.Println(student.Number, student.Name)
}
```

The plugin can be loaded by opening it through the `plugin` library in the main application.
The function that the plugin provides is found through the `Lookup` function.
It can then be executed by provided the signature and input.
The following code results in a greeting for the student defined in the plugin.

```go
helloPlugin, _ := plugin.Open("./plugins/hello.so")
sayHello, _ := helloPlugin.Lookup("SayHello")
sayHello.(func())()
```

A problem arises with this approach when trying to move data between the app and the plugin.
As long as one is dealing with primitive data types, this works fine.
Structures however, are not interchangeable between the app and the plugin.
A common library needs to be created that holds the definitions of the structures.
This library needs to be imported by both the app and the plugin.
Not the files, but the compiled versions of the library needs to be the same for both the broker and plugins.

```go
package common

type Student interface {
	GetName() string
	GetNumber() string
}
```

In the example above, a Student interface is defined in a common library.
Both can use this interface to exchange Student data.
Extended on this principle, a common plugin interface can be defined that all plugins need to implement.
This way, the app can load any plugin that implements this interface.

