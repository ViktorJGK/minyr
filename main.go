package main

import (
	"conv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	//src, err := os.Open("/home/viktor/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	outputFile, err := os.Create("kjevik-temp-fahrenheit-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}

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

		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			// Her
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				fahr := conv.CelsiusToFahrenheit(celsius)
				outputLine := fmt.Sprintf("%s;%s\n", celsius, fahr)
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
