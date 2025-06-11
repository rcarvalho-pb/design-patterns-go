package builderfluent

import (
	"fmt"
	"strings"
)

type person struct {
	name         string
	age          int
	workLocation string
	income       int
}

func (p person) GetName() string {
	return p.name
}

func (p person) GetAge() int {
	return p.age
}

func (p person) GetWorkLocation() string {
	return p.workLocation
}

func (p person) GetIncome() int {
	return p.income
}

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) SetWorkLocation(workLocation string) {
	p.workLocation = workLocation
}

func (p *person) SetIncome(income int) {
	p.income = income
}

func (p person) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Name: %s\n", p.GetName()))
	sb.WriteString(fmt.Sprintf("Age: %d\n", p.GetAge()))
	sb.WriteString(fmt.Sprintf("WorkPlace: %s\n", p.GetWorkLocation()))
	sb.WriteString(fmt.Sprintf("Income: %.2f\n", float64(p.GetIncome())/100.0))
	return sb.String()
}

type personBuilder struct {
	person *person
}

type personBuilderIdentity struct {
	personBuilder
}

type personBuilderWork struct {
	personBuilder
}

func NewPersonBuilder() *personBuilder {
	return &personBuilder{&person{}}
}

func (b *personBuilder) Identity() *personBuilderIdentity {
	return &personBuilderIdentity{*b}
}

func (b *personBuilder) Works() *personBuilderWork {
	return &personBuilderWork{*b}
}

func (b *personBuilderIdentity) WithName(name string) *personBuilderIdentity {
	b.person.name = name
	return b
}

func (b *personBuilderIdentity) WithAge(age int) *personBuilderIdentity {
	b.person.age = age
	return b
}

func (b *personBuilderWork) At(place string) *personBuilderWork {
	b.person.workLocation = place
	return b
}

func (b *personBuilderWork) Income(income int) *personBuilderWork {
	b.person.income = income
	return b
}

func (b *personBuilder) Build() *person {
	return b.person
}
