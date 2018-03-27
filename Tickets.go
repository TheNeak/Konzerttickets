package main


type Ticket struct {
	TicketId string `json: ticketID`
	Name     string `json:"name"`
	Date     string `json:"date"`
}

var tickets []Ticket

