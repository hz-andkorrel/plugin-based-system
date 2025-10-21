package domain

type BaseStudent struct {
	Name   string
	Number string
}

func (student *BaseStudent) GetName() string {
	return student.Name
}

func (student *BaseStudent) GetNumber() string {
	return student.Number
}
