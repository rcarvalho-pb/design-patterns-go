package main

import (
	"fmt"

	"github.com/rcarvalho-pb/design-pattern-go/builder/builderfluent"
	"github.com/rcarvalho-pb/design-pattern-go/builder/builderfunctional"
)

func main() {
	buildfluentExemple()
	fmt.Println("--------------------------------------")
	buildFunctionalExemple()
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
