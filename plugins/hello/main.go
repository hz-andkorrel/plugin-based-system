package main

import (
	"fmt"

	"plugins/common"
)

func SayHello(student common.Student) {
	fmt.Println(student.GetNumber(), student.GetName())
}
