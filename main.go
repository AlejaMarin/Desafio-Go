package main

import (
	"fmt"
	"log"

	"github.com/AlejaMarin/Desafio-Go/internal/tickets"
)

func main() {
	// Requerimiento 1
	total, err := tickets.GetTotalTickets("china")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
