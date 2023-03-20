package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Åpne csv filen
	inputFilename := "kjevik-temp-celsius-20220318-20230318.csv"
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("kunne ikke åpne filen %s: %s\n", inputFilename, err)
	}
	defer inputFile.Close()

	//lager CSV-leser og skriver
	reader := csv.NewReader(inputFile)
	writer := *csv.NewWriter(outputFile)

	// leser hver rad av inputfilen og konverter cel til fahr og lager ny fil med fahr
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("kunne ikke lese rad:", err)
			continue
		}

		//converterer celsius til fahrenheit
		celsius, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Println("ugyldig temperaturverdi", err)
			continue
		}

		/*
			CelsiusToFahrenheit
			record[5] = strconv.FormatFloat(fahrenheit, "f", 2, 64)
		*/

		//skriver til output filen
		err = writer.Write(record)
		if err != nil {
			fmt.Println("kunne ikke skrive rad:", err)
			continue
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("kunne ikke skrive til fil:", err)
		return
	}

	fmt.Println("Konveretringen fra celsius til fahrenheit er fulført")
}
