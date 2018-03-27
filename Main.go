package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	tickets = append(tickets, Ticket{TicketId: "1", Name: "ADTR", Date: "Tuesday"})
	tickets = append(tickets, Ticket{TicketId: "2", Name: "Volbeat", Date: "Monday"})
	tickets = append(tickets, Ticket{TicketId: "3", Name: "LP", Date: "Sunday"})
	log.Fatal(http.ListenAndServe(":8080", router))
}