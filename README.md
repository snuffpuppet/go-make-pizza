# go-make-pizza
Use Golang to make Pizzas concurrently

V1: A sequential pizza making robot

V2: A model with 3 stations and a pipeline between each. 2 robots are making pizza bases and one each doing sauce and topping

V3: an arbitrary number of cross skilled robots able to do any job pick up the next available job to be done

Since the constraint of making a pizza is that the base must be made then the sauce added and then the toppings, concurrency is limited but
it gets interesting if you have less robots than pizzas and the efficiency comes out in V3
