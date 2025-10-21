package main

import (
    "fmt"

    "plugins/class/domain"
)

func main() {
    mainClass := domain.Class{ Name: "HelloClass" }
    student := domain.ClassStudent{ Name: "Jimmaphy", Number: "S001", Class: &mainClass }
    fmt.Println(student.Class.Name)
}
