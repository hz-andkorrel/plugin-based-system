package main

import (
    "fmt"

    "plugins/class/domain"
)

func SayHello() {
    student := domain.ClassStudent{ Name: "Joseph", Number: "S001" }
    fmt.Println(student.Number, student.Name)
}

