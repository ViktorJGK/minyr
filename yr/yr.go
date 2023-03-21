package yr

import (
	"fmt"
	"strconv"
	"strings"

	_ "github.com/ViktorJGK/funtemps/conv"
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
		}

		celsius := conv.CelsiusToFahrenheit(lastPart)
		fmt.Printf("%.2f er det samme som %.2f i fahrenheit\n", lastPart, fahrenheit)

		// gjør at det er kunn 1 decimal plass
		lastPartStr := fmt.Sprintf("%.1f", lastPart)

		// kobler alle delene sammen igjen
		parts[3] = lastPartStr
		outputStr := strings.Join(parts, ";")

		//Forande slik at den gjør endringene i ny fil
		fmt.Println(outputStr)
	}
}

func convertionNull() {}

func convertionNegativ() {}

func convTextSlutt() {}
