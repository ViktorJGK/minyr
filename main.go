package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// åpner CSV filen
	fd, error := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Åpnet filen uten problemer")
	defer fd.Close()

	// Leser CSV filen
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(records)
}
