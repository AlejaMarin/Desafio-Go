package tickets

import (
	"errors"
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

func readData() []Ticket {
	//Open CSV file
	file, err := os.Open("tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Read CSV file
	rawData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(rawData), "\n")

	var tickets []Ticket
	for _, v := range data {

		line := strings.Split(v, ",")

		idTrim := strings.Trim(line[0],"\r")
		precioTrim := strings.Trim(line[5],"\r")

		id, err := strconv.ParseInt(idTrim, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		precio, err2 := strconv.ParseInt(precioTrim, 10, 32)
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
	file.Close()
	return tickets
}

// ejemplo 1
// func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
// func GetMornings(time string) (int, error) {}

// Requerimiento 3
// Calcular el porcentaje de personas que viajan a un país determinado en un día.
func PercentageDestination(destination string, total int, ch chan float64) (float64, error) {

	ts := readData()

	count := 0

	for _, v := range ts {
		if v.PaisDestino == destination {
			count++
		}
	}
	if count == 0 {
		return 0, errors.New("el destino ingresado es inválido")
	}

	porcentaje := (float64(count) / float64(total)) * 100
	ch <- porcentaje
	return porcentaje, nil
}
