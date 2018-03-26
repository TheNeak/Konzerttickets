package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func KonzertticketIndex(w http.ResponseWriter, r *http.Request) {
	Konzerttickets := Konzerttickets{
		Konzertticket{Name: "a day to remember"},
		Konzertticket{Name: "volbeat"},
	}

	if err := json.NewEncoder(w).Encode(Konzerttickets); err != nil {
		panic(err)
	}
}

func KonzertticketShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	KonzertticketId := vars["KonzertticketId"]
	fmt.Fprintln(w, "Konzertticket show:", KonzertticketId)
}