package yr

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ViktorJGK/funtemps/conv"
	"os"
	"strconv"
	"strings"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) == 4 {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil

	return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}

func CountLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
func AverageTemp(file *os.File) (float64, error) {
	scanner := bufio.NewScanner(file)
	var temperatures []float64
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Hopper over første og siste linje
		if lineCount == 0 || !scanner.Scan() {
			lineCount++
			continue
		}

		// Splitter linjene ved semikolon
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			continue
		}

		// Drar ut temperaturen
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			continue
		}
		temperatures = append(temperatures, temperature)

		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if len(temperatures) == 0 {
		return 0, fmt.Errorf("ingen valid temperatur oppdaget")
	}

	// Kalkulerer summen av hver femte element
	sum := 0.0
	for i := 4; i < len(temperatures); i += 5 {
		sum += temperatures[i]
	}

	// Kalkulerer gjennomsnitstemperaturen
	average := sum / float64(len(temperatures)/5)

	return average, nil
}
