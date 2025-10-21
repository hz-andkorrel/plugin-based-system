package main

import (
	"fmt"
	"plugins/base/domain"
)

func CreateDataStore() []domain.BaseStudent {
	return []domain.BaseStudent{
		{Name: "Alice", Number: "S001"},
		{Name: "Bob", Number: "S002"},
	}
}

func main() {
	students := CreateDataStore()

	fmt.Println("Student Records:")
	for _, student := range students {
		fmt.Printf("Name: %s, Number: %s\n", student.GetName(), student.GetNumber())
	}
}
