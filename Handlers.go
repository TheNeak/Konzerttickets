package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "log"
	"github.com/gorilla/mux"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo! Ticketübersicht bei /Tickets ... für bestimmtes Ticket /Tickets/ID")
}

func TicketIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tickets)
}

func TicketShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tickets {
		if item.TicketId == params["TicketId"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func TicketDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tickets {
		if item.TicketId == params["TicketId"] {
			tickets = append(tickets[:index], tickets[index+1:]...)
			log.Print(params["TicketId"])
			break
		}
		json.NewEncoder(w).Encode(tickets)
	}
}

func TicketPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)
	ticket.TicketId = params["ticketID"]
	tickets = append(tickets, ticket)
	json.NewEncoder(w).Encode(ticket)

}