package main

import "time"

type Ticket struct {
	TicketId  int       `json: TicketId`
	Name      string    `json:"name"`
	Due       time.Time `json:"due"`
}

type Tickets []Ticket