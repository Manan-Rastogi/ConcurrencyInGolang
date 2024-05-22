package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NoOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NoOfPizzas{
		delay := rand.Intn(5) + 1
		color.Blue(fmt.Sprintf("Received Order #%d!", pizzaNumber))

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd % 3 == 0{
			pizzaFailed++
		}else{
			pizzaMade++
		}

		color.Yellow("Making Pizza %d will take %d sec.", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd == 3 || rnd == 9 || rnd == 12{
			msg = fmt.Sprintf("*** We can out of ingredients while making pizza #%d", pizzaNumber)
		}else if rnd == 6{
			msg = fmt.Sprintf("*** Cook died while making pizza #%d", pizzaNumber)
		}else{
			success = true
			msg = fmt.Sprintf("Pizza #%d is Ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzafun(pizzaMaker *Producer){
	// keep track of which pizza we are making
	i := 0

	// run forever or until we receive a quit notification
	// try to make pizza
	for{
		// try to make a pizza
		// decision structure
		currentPizza := makePizza(i)
		if currentPizza != nil{
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			
			case quitChannel:= <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChannel)
				return
			}
		}
	}

}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Print out a message
	color.Cyan("Welcome To The Pizza Byte BCC Cafe")
	color.Cyan("__________________________________")

	// Create a Producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}


	// Run the producer in Background
	go pizzafun(pizzaJob)

	// Create and Run a consumer
	for i:= range pizzaJob.data{
		if i.pizzaNumber <= NoOfPizzas{
			if i.success{
				color.Green(i.message)
			}else{
				color.Red(i.message)
			}
		}else{
			color.Cyan("Done Making Pizzas")
			err := pizzaJob.Close()
			if err != nil{
				color.Red("Error Closing Channel: %v", err.Error())
			}
		}
	}

	// Print out the ending message
	color.Cyan("_______________")
	color.Cyan("Done for the day!!")

	color.Magenta(" ***** Summary for the Day: ******")
	color.HiMagenta("We made %d Pizzas and failed %d pizzas today.", pizzaMade, pizzaFailed)
}