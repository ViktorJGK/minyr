package main

import (
	"bufio"
	"fmt"
	"github.com/ViktorJGK/minyr/yr"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
			log.Printf("feil under skriving til outputfil: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Kopiering vellykket")

	// Oppdaterer outputfilen
	outputFile, err = os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// konverterer celsius til fahrenheit
	scanner = bufio.NewScanner(outputFile)
	var lines []string
	isFirstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine {
			lines = append(lines, line)
			isFirstLine = false
			continue
		}
		convertedLine, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			log.Printf("Feil for konvertering av Celsius til Fahrenheit i linje: '%s': %v\n", line, err)
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
			log.Printf("Feil under skriving til outputfil: %v\n", err)
		}
	}

	log.Println("Konversjon vellykket")

	// Leser outputfilen
	content, err := ioutil.ReadFile("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil ved lesing av fil:", err)
		return
	}

	// Konverterer innholdet til string
	contentStr := string(content)

	// Leter etter linje å erstatte
	index := strings.Index(contentStr, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;0.0")
	if index == -1 {
		fmt.Println("String ikke funnet i filen")
		return
	}

	// Erstatter stringen med ønsket string
	newContentStr := contentStr[:index] + "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Viktor J.G Kalhovd" + contentStr[index+len("Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;0.0"):]

	// Skriver stringen til outputfilen
	err = ioutil.WriteFile("kjevik-temp-fahr-20220318-20230318.csv", []byte(newContentStr), 0644)
	if err != nil {
		fmt.Println("Feil lesing til fil:", err)
		return
	}

	fmt.Println("String er byttet")
	fmt.Println("Filen kjevik celsius har blit konvertert til fahrenheit i den nye filen: kjevik-temp-fahr-20220318-20230318.csv")
}
