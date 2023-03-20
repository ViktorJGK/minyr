package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// definer flag-variablene

func init() {
	/*
					Implementer kommando i cmd:
			minyr
			convert
			average (ekstra oppgave)
			exit

		(se oppgave for detaljer)
	*/

}

func main() {
	// Åpne csv filen
	inputFilename := "kjevik-temp-celsius-20220318-20230318.csv"
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("kunne ikke åpne filen %s: %s\n", inputFilename, err)
	}
	defer inputFile.Close()

	// Lager outputfilen for skriving
	outputFilename := "kjevik-temp-fahr-20220318-20230318.csv"
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("kunne ikke oprette filen %s: %s\n", outputFilename, err)
	}

	//lager CSV-leser og skriver
	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(outputFile)

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
		cel, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Println("ugyldig temperaturverdi", err)
			continue
		}
		fahrenheit := (cel * 1.8) + 32.0
		record[4] = strconv.FormatFloat(fahrenheit, 'f', 2, 64)

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
