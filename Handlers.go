package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo! Ticketübersicht bei /Tickets ... für bestimmtes Ticket /Tickets/ID")
}

func TicketIndex(w http.ResponseWriter, r *http.Request) {
	Tickets := Tickets{
		Ticket{TicketId: 1, Name: "a day to remember"},
		Ticket{TicketId: 2, Name: "volbeat"},
	}

	if err := json.NewEncoder(w).Encode(Tickets); err != nil {
		panic(err)
	}
}

func TicketShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TicketId := vars["TicketId"]
	fmt.Fprintln(w, "Ticket show:", TicketId)
}

func TicketDelete(w http.ResponseWriter, r *http.Request) {


}