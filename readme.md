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


## Implementing go-plugings through an event bus

Another approach to creating a plugin system is through an event bus.
The architecture of the application is similar to the previous example.
The difference being the broker pattern being replaced by an event bus.
The event bus allows for communication between the app and the plugins through events.
Events are strings that are registered on the bus.
Plugins can subscribe to events and publish events.

Compared to the broker pattern, this approach is more flexible.
Plugins can subscribe to multiple events and publish events that other plugins can listen to.
This allows for a more dynamic interaction between the app and the plugins.
However, this approach also has its downsides.
The main issue is the lack of type safety.
Events are identified by strings, which can lead to errors if the event names are not consistent.
Additionally, logging and debugging can be more challenging due to the asynchronous nature of event-driven systems.

### Dockerized broker

This repository also contains a lightweight, Docker-friendly broker that treats plugins as independent HTTP services which self-register with the broker at POST /plugins/register. The dockerized subtree contains compose files, example plugin templates and an on-disk manifest (`dockerized/containers.yml`) that the broker reads and updates.

Quick start:

```bash
docker compose -f dockerized/compose.yaml up --build
```

Running plugins as containers has clear benefits: it removes language coupling, gives process isolation so a misbehaving plugin won't bring down the broker, and lets each plugin be deployed, scaled and observed independently with standard container tooling. At the same time this approach adds operational complexity — networking, health checks and keeping the plugin manifest durable — and introduces extra latency compared with in-process `.so` plugins.

In production this registration approach is fragile and potentially unsafe: duplicate or concurrent registrations can corrupt the manifest or leave the broker in an inconsistent state. Leaving the registration endpoint open is a serious security risk, and the simple demo proxy is not reliable enough for real traffic — it can drop or misroute requests and offers poor observability. Treat the demo components as examples only; they are not production-ready.
