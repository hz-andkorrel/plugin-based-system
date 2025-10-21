package domain

type BaseStudent struct {
	Name   string
	Number string
}

func NewBaseStudent(name, number string) *BaseStudent {
    return &BaseStudent{
        Name: name,
        Number: number,
    }
}

func (student *BaseStudent) GetName() string {
	return student.Name
}

func (student *BaseStudent) GetNumber() string {
	return student.Number
}
