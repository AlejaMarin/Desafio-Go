package tickets

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          int
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      int
}

// ejemplo 1
// func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
// func GetMornings(time string) (int, error) {}

// Requerimiento 3
// Calcular el porcentaje de personas que viajan a un país determinado en un día.
func PercentageDestination(destination string, total int) (float64, error) {

	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(rawData), "\n")

	var tickets []Ticket
	for _, v := range data {

		line := strings.Split(v, ",")

		id, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		precio, err2 := strconv.ParseInt(line[5], 10, 32)
		if err2 != nil {
			log.Fatal(err2)
		}

		t := Ticket{
			ID:          int(id),
			Nombre:      string(line[1]),
			Email:       string(line[2]),
			PaisDestino: string(line[3]),
			HoraVuelo:   string(line[4]),
			Precio:      int(precio)}
		tickets = append(tickets, t)

	}

	count := 0

	for _, v := range tickets {
		if v.PaisDestino == destination {
			count++
		} else {

			// Arrojar error, cuando el destino no exista o este mal escrito.
		}
	}

	porcentaje := (float64(count) / float64(total)) * 100
	return porcentaje, nil
}
