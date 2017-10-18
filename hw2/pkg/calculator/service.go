package calculator

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Greet greets a person
func Calculate(in string) string {
	if strings.Contains(in, "+") {
		calculation := strings.Split(in, "+")
		var calculationInt []float64
		for _, calculationStr := range calculation {
			i, err := strconv.ParseFloat(strings.Trim(calculationStr, " "), 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			calculationInt = append(calculationInt, i)
		}
		//fmt.Println(calculationInt[0] + calculationInt[1])
		return fmt.Sprintf("%f", calculationInt[0]+calculationInt[1])
	} else {
		calculation := strings.Split(in, "-")
		var calculationInt []float64

		minusBefore := false
		for _, calculationStr := range calculation {
			if calculationStr == "" {
				minusBefore = !minusBefore
			} else {
				i, err := strconv.ParseFloat(strings.Trim(calculationStr, " "), 64)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				if minusBefore {
					i *= -1
					minusBefore = false
				}
				calculationInt = append(calculationInt, i)
			}
		}
		return fmt.Sprintf("%f", calculationInt[0]-calculationInt[1])
	}
}
