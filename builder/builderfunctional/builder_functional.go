package builderfunctional

import "fmt"

type person struct {
	name      string
	age       int
	workplace string
	income    int
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nWorkPlace: %s\nIncome: %.2f\n",
		p.name,
		p.age,
		p.workplace,
		float64(p.income)/100.0,
	)
}

type modPerson func(*person)
type personBuilderFunctional struct {
	actions []modPerson
}

func NewPersonBuilderFunctional() *personBuilderFunctional {
	return &personBuilderFunctional{make([]modPerson, 0)}
}

func (b *personBuilderFunctional) Build() *person {
	p := &person{}
	for _, action := range b.actions {
		action(p)
	}
	return p
}

func (b *personBuilderFunctional) WithName(name string) *personBuilderFunctional {
	b.actions = append(b.actions, func(p *person) {
		p.name = name
	})
	return b
}

func (b *personBuilderFunctional) WithAge(age int) *personBuilderFunctional {
	b.actions = append(b.actions, func(p *person) {
		p.age = age
	})
	return b
}

func (b *personBuilderFunctional) WithWorkPlace(workplace string) *personBuilderFunctional {
	b.actions = append(b.actions, func(p *person) {
		p.workplace = workplace
	})
	return b
}

func (b *personBuilderFunctional) WithIncome(income int) *personBuilderFunctional {
	b.actions = append(b.actions, func(p *person) {
		p.income = income
	})
	return b
}
