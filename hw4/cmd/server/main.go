package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"gitlab.3fs.si/mdi/go-academy-hw/hw4/pkg/bucket"

	"gitlab.3fs.si/mdi/go-academy-hw/hw4/pkg/calculator"
)

var dbReadWriter bucket.MockReadWriter
var currentId int = 1

func addToDB(result string) {
	resultInt, _ := strconv.ParseInt(result, 0, 0)
	dbReadWriter.Insert(currentId, int(resultInt))
	currentId++
}

func getValuesFromDb(ids []int) ([]int, error) {
	var idValues []int = make([]int, len(ids))
	for idx, id := range ids {
		val, err := dbReadWriter.Get(id)
		if err != nil {
			return nil, err
		} else {
			idValues[idx] = val
		}
	}
	return idValues, nil
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<form method="post">
		<input name="calculation" required> <input type="submit" value="Calculate expression">
		</form>
		{{ if .Calculator }}<h1>{{ .Calculator }}</h1>{{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	calculation := ""
	if r.FormValue("calculation") != "" {
		calculatedVal, idsFromDB, operator := calculator.Parse(r.FormValue("calculation"))
		if idsFromDB == nil || len(idsFromDB) == 0 {
			addToDB(calculatedVal)
			calculation += calculatedVal
		} else {
			valuesFromDB, err := getValuesFromDb(idsFromDB)
			if err != nil {
				calculation = err.Error()
			} else {
				finalValue := calculator.CalculateArrays(calculatedVal, valuesFromDB, operator)
				calculation += finalValue
				addToDB(finalValue)
			}
		}
	}

	// prepare the data
	data := struct {
		Calculator string
	}{
		Calculator: calculation,
	}

	// parse the template
	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		fmt.Println("Failed to parse template;", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func startServer(address string) {
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
	var addr = flag.String("addr", ":8080", "Interface and port to listen on")
	err := errors.New("Key not found")
	dbReadWriter, err = bucket.CreateMockDB("MockDB")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// parse the flags
	flag.Parse()

	startServer(*addr)
}
