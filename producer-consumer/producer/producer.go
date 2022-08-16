package producer

import (
	"github.com/adamqazi/go-examples/producer-consumer/ticket"
)

// Producer is responsible for routing generated tickets.
type Producer struct {
	Data chan ticket.Ticket
	Quit chan chan error
}

// NewProducer creates a new producer.
func NewProducer() *Producer {
	return &Producer{
		Data: make(chan ticket.Ticket),
		Quit: make(chan chan error),
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.Quit <- ch

	return <-ch
}
