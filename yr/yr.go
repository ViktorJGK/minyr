package yr

import (
	"strconv"
	"strings"

	"github.com/ViktorJGK/funtemps/conv"
)

func convertionPositiv() {
	inputArr := []string{}
	for i := range inputArr {
		inputStr := inputArr[i]

		//deler innput i strengen mellom ;
		parts := strings.Split(inputStr, ";")

		lastPart, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			panic(err)
			conv.CelsiusToFahrenheit(cel)

		}
	}
}

func convertionNull() {}

func convertionNegativ() {}
