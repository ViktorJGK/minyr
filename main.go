package main

import (
	"bufio"
	"fmt"
	"github.com/ViktorJGK/minyr/yr"
	"log"
	"os"
)

func main() {

	// Åpner src filen
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// Lager output filen
	outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Leser data fra src filen og skriver til output filen
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		outputLine := fmt.Sprintf("%s\n", line)
		_, err = outputFile.WriteString(outputLine)
		if err != nil {
			log.Printf("Error writing to output file: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Copying complete. Results written to kjevik-temp-fahr-20220318-20230318.csv")

	// Update the last element in the output file
	outputFile, err = os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner = bufio.NewScanner(outputFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		convertedLine, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			log.Printf("Error converting Celsius to Fahrenheit for line '%s': %v\n", line, err)
			continue
		}
		lines = append(lines, convertedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	outputFile.Truncate(0)
	outputFile.Seek(0, 0)

	for _, line := range lines {
		outputLine := fmt.Sprintf("%s\n", line)
		_, err = outputFile.WriteString(outputLine)
		if err != nil {
			log.Printf("Error writing to output file: %v\n", err)
		}
	}

	log.Println("Conversion complete. Results written to kjevik-temp-fahr-20220318-20230318.csv")
}
