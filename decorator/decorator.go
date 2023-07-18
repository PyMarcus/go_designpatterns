package main

import "fmt"

// Decorator is a structural design pattern that
// lets attach new behaviors to objects.
type IPizza interface {
	getPrice() float64
}

type Pizza struct{
	price float64
}

func (p Pizza) getPrice() float64{
	return p.price
}

type PizzaTomato struct {
	pizza IPizza
}

func (p PizzaTomato) getPrice() float64{
	fmt.Print("Tomato pizza - ")
	return p.pizza.getPrice()
}

func main() {

	// decore generical pizza
	pizza := &Pizza{33.80} 

	pizzaTomato := PizzaTomato{pizza: pizza}

	fmt.Println("Pizza: ", pizzaTomato.getPrice())
}
