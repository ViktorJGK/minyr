package main

import (
	"bufio"
	"fmt"
	"github.com/ViktorJGK/minyr/yr"
	"log"
	"os"
)

func main() {

	// åpner orginal filen
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	//src, err := os.Open("/home/viktor/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// lager output filen
	outputFile, err := os.Create("kjevik-temp-fahrenheit-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	//Leser data fra kilde filen til og skriver til den nye filen
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

	log.Println("Kopiering vellykket. Results written to kjevik-temp-fahr-20220318-20230318.csv")

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

	log.Println("konversjon vellykket")
}

/*
	log.Println(src)

	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
	bytesCount := 0
	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}



/*
		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			// Her
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsiusStr := elementArray[3]
				celsius, err := strconv.ParseFloat(celsiusStr, 64)
				if err != nil {
					log.Printf("Feil ved konvertering celsius til float: %v\n", err)
					continue
				}
				fahr := conv.CelsiusToFahrenheit(celsius)
				outputLine := fmt.Sprintf("%s;%s\n", celsiusStr, strconv.FormatFloat(fahr, 'f', -1, 64))
				outputFile.Write([]byte(outputLine))
			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}
}
*/
