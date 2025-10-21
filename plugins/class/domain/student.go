package domain

type ClassStudent struct {
	Name   string
	Number string
    Class  *Class
}

func (student *ClassStudent) GetName() string {
	return student.Name
}

func (student *ClassStudent) GetNumber() string {
	return student.Number
}
