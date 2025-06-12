package main

import (
	"fmt"

	"github.com/rcarvalho-pb/design-pattern-go/builder/builderfluent"
	"github.com/rcarvalho-pb/design-pattern-go/builder/builderfunctional"
	"github.com/rcarvalho-pb/design-pattern-go/factory/factorybuilder"
	"github.com/rcarvalho-pb/design-pattern-go/factory/factoryfunctional"
)

func separation() {
	fmt.Println("--------------------------------------")
}

func main() {
	buildfluentExemple()
	separation()
	buildFunctionalExemple()
	separation()
	factoryFunctionalExemple()
	separation()
	factoryBuilderExemple()
}

func buildFunctionalExemple() {
	ramon := builderfunctional.NewPersonBuilderFunctional().
		WithName("Ramon").
		WithAge(30).
		WithWorkPlace("Energisa").
		WithIncome(3900).
		Build()
	fmt.Println(ramon)
}

func buildfluentExemple() {
	ramon := builderfluent.NewPersonBuilder().
		Identity().
		WithName("Ramon").
		WithAge(30).
		Works().
		At("Energisa").
		Income(3900).
		Build()
	fmt.Println(ramon)
}

func factoryFunctionalExemple() {
	engeneeringFactory := factoryfunctional.NewEmployeeFactory("Developer", 60000)
	ramon := engeneeringFactory("Ramon")
	fmt.Println(ramon)
}

func factoryBuilderExemple() {
	developerBuilder := factorybuilder.NewEmployeeBuilder("Developer", 60000)
	ramon := developerBuilder.Create("Ramon")
	fmt.Println(ramon)
	developerBuilder.AnualIncome = 80000
	emilly := developerBuilder.Create("Emilly")
	fmt.Println(emilly)
}
