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

	// konverterer celsius til fahrenheit
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

	// Read the contents of file2
	content, err := ioutil.ReadFile("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the content to a string
	contentStr := string(content)

	// Find the index of the string to be replaced
	index := strings.Index(contentStr, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;0.0")
	if index == -1 {
		fmt.Println("String not found in file")
		return
	}

	// Replace the string with the desired text
	newContentStr := contentStr[:index] + "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Viktor JG Kalhovd" + contentStr[index+len("Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;0.0"):]

	// Write the updated content back to file2
	err = ioutil.WriteFile("kjevik-temp-fahr-20220318-20230318.csv", []byte(newContentStr), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("String replaced successfully!")
}
