# go-make-pizza
Use Golang to make Pizzas concurrently

V1: A sequential pizza making robot - no concurrency just a simple program

V2: A model with 3 stations and a pipeline between each. 2 robots are making pizza bases and one each doing sauce and topping
  - Channels are used to transfer activity between each pizza making station and routines represent robots working at each station
  
V3: an arbitrary number of cross skilled robots able to do any job pick up the next available job to be done
  - A single channel holds all the requests for work on the Pizza. a number of robots listen to the channel and 
    pick up jobs to do and return a status code to the requestor
    Each requestor ensures the correct sequence of Pizza making activities by sequentially requesting those jobs to be done
    
Since the constraint of making a pizza is that the base must be made then the sauce added and then the toppings, concurrency is limited but
it gets interesting if you have less robots than pizzas and the efficiency comes out in V3
