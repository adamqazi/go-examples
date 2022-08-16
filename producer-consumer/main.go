package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/adamqazi/go-examples/producer-consumer/producer"
	"github.com/adamqazi/go-examples/producer-consumer/ticket"
	"github.com/fatih/color"
)

const (
	totalTickets = 25
)

var (
	ticketsCreated   = 0
	ticketsProcessed = 0
)

func generateTicket() *ticket.Ticket {
	ticketsCreated++
	usr := fmt.Sprintf("user-%v", ticketsCreated)
	msg := "Please resolve issue ASAP!"

	color.Yellow("generating ticket for user %v\n", usr)

	delay := time.Duration(rand.Int63n(5)) * time.Second
	time.Sleep(delay)

	return ticket.NewTicket(usr, msg)
}

func runProducer(prd *producer.Producer) {
	for {
		ticket := generateTicket()
		if ticket != nil {
			select {
			case prd.Data <- *ticket:
			case quitChan := <-prd.Quit:
				close(prd.Data)
				close(quitChan)

				return
			}
		}
	}
}

func main() {
	color.Cyan("producer-consumer problem")
	color.Cyan("-------------------------")

	rand.Seed(time.Now().UnixNano())

	// create a producer.
	job := producer.NewProducer()

	// run the producer in the background.
	go runProducer(job)

	// create and run consumer.
	for tkt := range job.Data {
		if ticketsProcessed > totalTickets {
			color.Red("all tickets processed, closing channels")

			err := job.Close()
			if err != nil {
				color.Red("unable to close channels, %v", err)
			}
		} else {
			tkt.ResolveTicket()
			ticketsProcessed++
		}
	}

	color.Cyan("all tickets resolved")
	color.Cyan("--------------------")
}
