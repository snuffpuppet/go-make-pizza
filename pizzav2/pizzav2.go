package main

import "fmt"
import "time"
import "strconv"
import "os"

func MakeDough(ich chan string, och chan string) {
	for pizza := range ich {
		fmt.Printf("Making dough for %s\n", pizza)
		time.Sleep(time.Millisecond*5000)
		och <- pizza
	}
}

func AddSauce(ich chan string, och chan string) {
	for pizza := range ich {
		fmt.Printf("Adding sauce for %s\n", pizza)
		time.Sleep(time.Millisecond*2000)
		och <- pizza
	}
}

func AddToppings(ich chan string, och chan string) {
	for pizza := range ich {
		fmt.Printf("Adding toppings for %s\n", pizza)
		time.Sleep(time.Millisecond*2000)
		och <- pizza
	}
}

func MakePizza(num int) {
	fmt.Printf("Makeing %d Pizza(s)\n", num)

	doughch := make(chan string, num)
	saucech := make(chan string, num)
	toppch := make(chan string, num)
	donech := make(chan string, num)

	go MakeDough(doughch, saucech)
	go MakeDough(doughch, saucech)
	go AddSauce(saucech, toppch)
	go AddToppings(toppch, donech)
	
	for i:=0; i<num; i++ {
		pizza := "Pizza #" + strconv.Itoa(i);
		doughch <- pizza
	}

	for i := 0; i<num; i++ {
		pizza := <- donech
		fmt.Printf("Done %s\n", pizza);
	}
}

func main() {
	numPizzas:=1
	
	if (len(os.Args) == 2) {
		if num,err:=strconv.Atoi(os.Args[1]); err == nil {
			numPizzas=num
		}
	}
	MakePizza(numPizzas)
}
