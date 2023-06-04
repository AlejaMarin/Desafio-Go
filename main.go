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

	// Requerimiento 1:
	// total, err := tickets.GetTotalTickets("Brazil")

	// Requerimiento 2:
	cantidad, timeTravel, f := tickets.GetCountByPeriod("tard")
	if f != nil {
		log.Fatal(f)
	}
	fmt.Println("La cantidad de personas que viajan por la", timeTravel, "es de:", cantidad)

	// Requerimiento 3:
	/*total, err := tickets.PercentageDestination("Brazil", 1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Porcentaje Personas:", total, "%")*/

}
