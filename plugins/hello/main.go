package main

import (
    "fmt"

    "plugins/class/common"
)

func SayHello(student common.Student) {
    fmt.Println(student.GetNumber(), student.GetName())
}

