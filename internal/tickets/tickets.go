package tickets

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

		idTrim := strings.Trim(line[0], "\r")
		precioTrim := strings.Trim(line[5], "\r")

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

// Requerimiento 1
// Una función que calcule cuántas personas viajan a un país determinado.
func GetTotalTickets(destination string, can chan int) (int, error) {

	//Open CSV file
	file, err := os.Open("tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Read CSV file
	rowData, err := os.ReadFile("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(rowData), "\n")

	total := 0

	for _, value := range data {
		line := strings.Split(value, ",")
		pais := line[3]

		if strings.EqualFold(pais, destination) {
			total++
		}
	}

	if total == 0 {
		return 0, fmt.Errorf("pais no encontrado: %s", destination)
	}

	can <- total
	return total, nil

}

// Requerimiento 2
/* Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6),
mañana (7 → 12), tarde (13 → 19), y noche (20 → 23). */
func GetCountByPeriod(flightTime string) (int, string, error) {

	ts := readData()

	countDawning := 0
	countMorning := 0
	countAfternoon := 0
	countNight := 0

	for _, w := range ts {

		t, err := time.Parse("15:04", w.HoraVuelo)
		if err != nil {
			log.Fatal(err)
		}

		if t.Hour() >= 0 && t.Hour() <= 6 {
			countDawning++
		}

		if t.Hour() >= 7 && t.Hour() <= 12 {
			countMorning++
		}

		if t.Hour() >= 13 && t.Hour() <= 19 {
			countAfternoon++
		}

		if t.Hour() >= 20 && t.Hour() <= 23 {
			countNight++
		}

	}

	switch flightTime {
	case "madrugada":
		return countDawning, flightTime, nil
	case "maniana":
		return countMorning, flightTime, nil
	case "tarde":
		return countAfternoon, flightTime, nil
	case "noche":
		return countNight, flightTime, nil
	default:
		return 0, "", fmt.Errorf("el valor ingresado %s es inválido", flightTime)
	}
}

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
