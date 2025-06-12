package factoryfunctional

import "fmt"

type Employee struct {
	name, position string
	anualIncome    int
}

func (e *Employee) String() string {
	return fmt.Sprintf("Name: %s\nPosition: %s\nAnual Income: %.2f", e.name, e.position, float64(e.anualIncome)/100.0)
}

func NewEmployeeFactory(position string, anualIncome int) func(string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, anualIncome}
	}
}
