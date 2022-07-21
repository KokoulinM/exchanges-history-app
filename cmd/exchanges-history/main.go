package main

import (
	"fmt"
	"log"

	"github.com/KokoulinM/exchanges-history-app/internal/csv"
)

func main() {
	data, err := csv.Reader("history.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(data))
}
