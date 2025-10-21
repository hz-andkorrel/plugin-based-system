# Research: Creating a plugin based system

This repository holds the demo's for several implementations of plugin systems.
The findings will be discussed in this document.
Each heading describes a single solution.


## Directly implementing go-plugins

Go offers a build-in solution for plug-ins.
By supply the `-buildmode=plugin` flag, the compiler will create a `.so` file.
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
Once struct starts to play, this functionality breaks.
An attempt was made to create a common library that holds the interface for a student.
But even with a fully mathing interface, the new execution will error.
Providing a `plugin.Symbol is func(domain.ClassStudent), not func(domain.BaseStudent)` error.
Even though both are using the common library interface which is shown below.

```go
package common

type Student interface {
	GetName() string
	GetNumber() string
}
```

**Concluding,** one could say the method of directly invoking plug-ins is not a suitable solution.
Even if the common library solution had been useful, version compatibility would have been an issue down the line.
Especially if the end goal would be to have external developers create their own plugins.
