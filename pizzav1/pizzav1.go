package main

import "fmt"
import "time"
import "strconv"

func MakeDough(pizza string) {
	fmt.Printf("Making dough for %s\n", pizza)
	time.Sleep(time.Millisecond*10000)
}

func AddSauce(pizza string) {
	fmt.Printf("Adding sauce for %s\n", pizza)
	time.Sleep(time.Millisecond*5000)
}

func AddToppings(pizza string) {
	fmt.Printf("Adding toppings for %s\n", pizza)
	time.Sleep(time.Millisecond*5000)
}

func Ship(pizza string) {
	fmt.Printf("Shipping %s\n", pizza)
}

func MakePizza(num int) {
	for i:=0; i<num; i++ {
		pizza := "Pizza #" + strconv.Itoa(i);
		MakeDough(pizza)
		AddSauce(pizza)
		AddToppings(pizza)
		Ship(pizza)
	}
}

func main() {
	MakePizza(1)
}
