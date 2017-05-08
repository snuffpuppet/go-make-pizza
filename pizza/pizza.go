package main

import "fmt"
import "time"
import "strconv"
import "os"

type Worker struct {
	ch chan Request
	fn func(pizza string)
}

type Request struct {
	fn func(pizza string) int
	pizza string 
	ch chan int
}

func MakeDough(pizza string) int {
	fmt.Printf("Making dough for %s\n", pizza)
	time.Sleep(time.Millisecond*5000)
	return 0
}

func AddSauce(pizza string) int {
	fmt.Printf("Adding sauce for %s\n", pizza)
	time.Sleep(time.Millisecond*2000)
	return 0
}

func AddToppings(pizza string) int {
	fmt.Printf("Adding toppings for %s\n", pizza)
	time.Sleep(time.Millisecond*2000)
	return 0
}

func requester(p string, work chan<- Request, done chan string) {
	c := make(chan int)
        work <- Request{MakeDough, p, c} // send request
        _ = <-c              // wait for answer
        work <- Request{AddSauce, p, c} // send request
        _ = <-c              // wait for answer
        work <- Request{AddToppings, p, c} // send request
        _ = <-c              // wait for answer
	done <- p
}

func DoWork(ch chan Request) {
	for {
		pizza := <- ch
		pizza.ch <- pizza.fn(pizza.pizza)
	}
}

func MakePizza(numP int, numW int) {
	fmt.Printf("Making %d Pizza(s) with %d workers\n", numP, numW)

	ch := make(chan Request, numW)
	donech := make(chan string, numP)

	// asynchonously request all the pizzas to be made
	for i:=0; i<numP; i++ {
		pizza := "Pizza #" + strconv.Itoa(i)
		go requester(pizza,ch, donech)
	}

	// Assign a load of workers to the channel
	for i:=0; i<numW; i++ {
		go DoWork(ch)
	}

	for i := 0; i<numP; i++ {
		pizza := <- donech
		fmt.Printf("Done %s\n", pizza);
	}

}

func main() {
	numPizzas:=1
	numWorkers:=5
	
	if (len(os.Args) >= 2) {
		if num,err:=strconv.Atoi(os.Args[1]); err == nil {
			numPizzas=num
		}
		if (len(os.Args) == 3) {
			if num,err:=strconv.Atoi(os.Args[2]); err == nil {
				numWorkers=num
			}
		}
	}
	MakePizza(numPizzas, numWorkers)
}
