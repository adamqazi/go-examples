package ticket

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

// Ticket represents an issue reported by a customer.
type Ticket struct {
	id         string
	message    string
	reportedBy string
	reportedAt time.Time
	resolved   bool
	resolvedAt time.Time
}

// NewTicket creates a new ticket with the provided message and reporter.
func NewTicket(usr, msg string) *Ticket {
	return &Ticket{
		id:         uuid.New().String(),
		message:    msg,
		reportedBy: usr,
		reportedAt: time.Now().UTC(),
		resolved:   false,
		resolvedAt: time.Time{},
	}
}

// ResolveTicket marks the ticket as resolved.
func (t *Ticket) ResolveTicket() {
	t.resolvedAt = time.Now().UTC()
	t.resolved = true

	fmt.Println(t.String())
	color.Green("ticket resolved")
	fmt.Println()
}

func (t *Ticket) String() string {
	return fmt.Sprintf("ID: %v\nMessage: %v\nReported By: %v\nReported At: %v\nResolved: %v\nResolved At: %v",
		t.id, t.message, t.reportedBy, t.reportedAt, t.resolved, t.resolvedAt)
}
