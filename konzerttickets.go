package main

import "time"

type Konzertticket struct {
	Name      string    `json:"name"`
	Due       time.Time `json:"due"`
}

type Konzerttickets []Konzertticket