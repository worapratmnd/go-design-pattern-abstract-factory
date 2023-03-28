package main

import (
	"example/abstract/factory/models"
	"fmt"
)

func main() {
	adidasFactory, _ := models.GetSportFactory("adidas")
	nikeFactory, _ := models.GetSportFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	nikeShoe := nikeFactory.MakeShoe()

	adidasShirt := adidasFactory.MakeShirt()
	nikeShirt := nikeFactory.MakeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)

}

func printShoeDetails(s models.IShoe) {
	fmt.Println("Shoes")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s models.IShirt) {
	fmt.Println("Shir")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
