package main

import (
	"fmt"

	"github.com/AlejaMarin/Desafio-Go/internal/tickets"
)

func main() {

	// Requerimiento 4:
	/* Ejecutar al menos una vez cada requerimiento en la función main.
	Las ejecuciones deben realizarse de manera concurrente (utilizando diferentes goroutines). */

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

	// Requerimiento 3
	ch3 := make(chan float64)
	go tickets.PercentageDestination("Colombia", 1000, ch3)
	fmt.Printf("El porcentaje de personas que viajan al pais ingresado es de: %.2f", <-ch3)

	/*total, err := tickets.PercentageDestination("Brazil", 1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Porcentaje Personas:", total, "%")*/

}
