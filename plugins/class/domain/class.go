package domain

type Class struct {
    Name string
}

func (class *Class) GetName() string {
    return class.Name
}
