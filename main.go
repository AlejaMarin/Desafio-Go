package main

import (
	"fmt"
	"log"

	"github.com/AlejaMarin/Desafio-Go/internal/tickets"
)

func main() {

	// Requerimiento 4:
	/* Ejecutar al menos una vez cada requerimiento en la función main.
	Las ejecuciones deben realizarse de manera concurrente (utilizando diferentes goroutines). */

	// total, err := tickets.GetTotalTickets("Brazil")

	total, err := tickets.PercentageDestination("Brazil", 1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Porcentaje Personas:", total, "%")

}
