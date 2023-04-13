package main

import (
	"bufio"
	"fmt"
	"github.com/ViktorJGK/minyr/yr"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("q/exit for å gå ut")
	fmt.Println("convert: for å konvertere filen")
	fmt.Println("average: for å få gjennomsnitsstemperatur")
	fmt.Println("")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			filePath := "kjevik-temp-fahr-20220318-20230318.csv"
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				fmt.Println("feil", err)
				fmt.Println("")
			}

			if fileInfo != nil {
				fmt.Println("fil eksisterer ønsker du å generere den på nytt? j/n")
				scanner.Scan()
				input = scanner.Text()
				if input == "j" {
					fmt.Println("genererer filen på nytt")
					konvert()
					fmt.Println("filen er generert på nytt")
					fmt.Println("")
					fmt.Println("q/exit for å gå ut")
					fmt.Println("convert: for å konvertere filen")
					fmt.Println("average: for å få gjennomsnitsstemperatur")
					fmt.Println("")

				} else if input == "n" {
					fmt.Println("gjør ingenting med filen")
					fmt.Println("")
					fmt.Println("q/exit for å gå ut")
					fmt.Println("convert: for å konvertere filen")
					fmt.Println("average: for å få gjennomsnitsstemperatur")
					fmt.Println("")
				}
			} else {
				fmt.Println("Konverterer alle målingene gitt i grader Celsius til grade Fahrenheit.")
				fmt.Println("")
				konvert()
			}

		}

		if input == "average" {
			fmt.Println("vil du se gjennomsnitt temperetur for celsius eller fahrenheit?")
			fmt.Println("c: for celsius")
			fmt.Println("f: for fahrenheit")
			fmt.Println("")

			scanner.Scan()
			input = scanner.Text()
			if input == "c" {
				fmt.Println("Gjennomsnittstemperatur i Celsius er følgende:")
				cAverage()

			} else if input == "f" {
				fmt.Println("Gjennom snitt temperatur i Fahrenheit er følgene:")
				fAverage()
			}
		}
	}
}

func konvert() {
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
	log.Println("String er byttet")
	log.Println("Filen kjevik celsius har blit konvertert til fahrenheit i den nye filen: kjevik-temp-fahr-20220318-20230318.csv")
}

func cAverage() {
	// Åpner cel filen
	cel, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer cel.Close()

	average, err := yr.AverageTemp(cel)
	if err != nil {
		log.Printf("Feil under gjennomsnittstemperatur måling: %v\n", err)
		return
	}
	fmt.Printf("Average Temperature: %.2f\n", average)
}
func fAverage() {
	// Åpner fahr filen
	fahr, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fahr.Close()

	average, err := yr.AverageTemp(fahr)
	if err != nil {
		log.Printf("Feil under gjennomsnittstemperatur måling: %v\n", err)
		return
	}
	fmt.Printf("Average Temperature: %.2f\n", average)
}

func minyrKjør() {
	// Kompilerer main.go filen ved hjelp av 'go build' kommandoen
	cmd := exec.Command("go", "build", "main.go")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Kompilering feilet:", err)
		return
	}

	// Sjekker om den kompilerte filen eksisterer
	if _, err := os.Stat("./main"); os.IsNotExist(err) {
		fmt.Println("Kompilert fil ble ikke funnet")
		return
	}

	// Setter tillatelse til å kjøre på den kompilerte filen
	cmd = exec.Command("chmod", "+x", "./main")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Feil ved setting av kjøretillatelse:", err)
		return
	}

	// Kjører det kompilerte programmet
	cmd = exec.Command("./main")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Kjøring av programmet feilet:", err)
		return
	}

	// Eventuell annen logikk eller handlinger etter kjøring av programmet
	fmt.Println("Programmet ble kjørt suksessfullt")
}
