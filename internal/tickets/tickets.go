package tickets

import (
	
	"fmt"
	"log"
	"os"
	"strings"
)

// ejemplo 1

func GetTotalTickets(destination string) (int, error) {
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

	return total, nil
	
}

// ejemplo 2
//func GetMornings(time string) (int, error) {}

// ejemplo 3
//func AverageDestination(destination string, total int) (int, error) {}
