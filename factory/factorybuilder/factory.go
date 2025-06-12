package factorybuilder

import "fmt"

type employee struct {
	name, position string
	anualIncome    int
}

func (e *employee) String() string {
	return fmt.Sprintf("Name: %s\nPosition: %s\nAnual Income: %.2f", e.name, e.position, float64(e.anualIncome)/100.0)
}

type employeeBuilder struct {
	Position    string
	AnualIncome int
}

func NewEmployeeBuilder(position string, anualIncome int) *employeeBuilder {
	return &employeeBuilder{position, anualIncome}
}

func (b *employeeBuilder) Create(name string) *employee {
	return &employee{name, b.Position, b.AnualIncome}
}
