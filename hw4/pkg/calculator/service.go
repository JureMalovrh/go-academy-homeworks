package calculator

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculate expression
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

func parseID(id string, err error) int {
	if strings.Contains(id, "$") {
		idx, err := strconv.ParseInt(strings.Trim(id, "$"), 0, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return int(idx)
	}
	fmt.Println(err)
	os.Exit(1)
	return 0
}

// Parse Calculation expression, returns array of strings to calculate or
func Parse(in string) (string, []int, string) {
	if strings.Contains(in, "+") {
		calculation := strings.Split(in, "+")
		var calculationInt []int
		var calculationIds []int
		for _, calculationStr := range calculation {
			i, err := strconv.ParseInt(strings.Trim(calculationStr, " "), 0, 0)
			if err != nil {
				idx := parseID(calculationStr, err)
				calculationIds = append(calculationIds, idx)
			} else {
				calculationInt = append(calculationInt, int(i))
			}
		}
		//if both numbers are actually numbers, return sumation
		if len(calculationInt) == 2 {
			calculatedStr := fmt.Sprintf("%d", calculationInt[0]+calculationInt[1])
			return calculatedStr, nil, ""
		}
		// else return string of first (len 1 or 0) and ids
		calculatedStr := ""
		for _, calcStr := range calculationInt {
			calculatedStr += fmt.Sprintf("%d", calcStr)
		}
		return calculatedStr, calculationIds, "+"

	} else {
		calculation := strings.Split(in, "-")
		var calculationInt []int
		var calculationIds []int

		minusBefore := false
		for _, calculationStr := range calculation {
			if calculationStr == "" {
				minusBefore = !minusBefore
			} else {
				i, err := strconv.ParseInt(strings.Trim(calculationStr, " "), 0, 0)
				if err != nil {
					idx := parseID(calculationStr, err)
					calculationIds = append(calculationIds, idx)
				} else {
					if minusBefore {
						i *= -1
						minusBefore = false
					}
					calculationInt = append(calculationInt, int(i))
				}
			}
		}
		if len(calculationInt) == 2 {
			calculatedStr := fmt.Sprintf("%d", calculationInt[0]-calculationInt[1])
			return calculatedStr, nil, ""
		}
		// else return string of first (len 1 or 0) and ids
		calculatedStr := ""
		for _, calcStr := range calculationInt {
			calculatedStr += fmt.Sprintf("%d", calcStr)
		}
		return calculatedStr, calculationIds, "-"
	}
}

// CalculateArrays obtained from Parse
func CalculateArrays(value1 string, valuesFromIds []int, operator string) string {
	var finalValue int
	if value1 != "" {
		value1int, _ := strconv.ParseInt(value1, 0, 0)
		if operator == "+" {
			finalValue += int(value1int)
		} else {
			finalValue -= int(value1int)
		}
	}

	for _, valueFromID := range valuesFromIds {
		finalValue += valueFromID
	}
	return fmt.Sprintf("%d", finalValue)
}
