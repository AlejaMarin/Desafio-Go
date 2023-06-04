package main

import (
	"fmt"

	"github.com/AlejaMarin/Desafio-Go/internal/tickets"
)

func main() {

	// Requerimiento 1
	/* go func() {

		men, total, err := tickets.GetTotalTickets("china")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("En el destino %s viajan %v personas", men, total)

	}()

	time.Sleep(time.Second * 2) */

	ch1 := make(chan int)
	go tickets.GetTotalTickets("china", ch1)
	fmt.Println("En el destino ingresado viajan", <-ch1)

}
