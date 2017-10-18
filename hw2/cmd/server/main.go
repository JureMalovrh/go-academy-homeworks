package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"gitlab.3fs.si/mdi/go-academy-hw/hw2/pkg/calculator"
	"gitlab.3fs.si/mdi/go-academy-hw/hw2/pkg/greeter"
)

type apiHandler struct{}

func (*apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// read and validate the query string
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// greet the user
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<form method="post">
		<input name="name" required> <input type="submit" value="Greet!">
		</form>
		{{ if .Greet }}<h1>{{ .Greet }}</h1>{{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	greet := ""
	if r.FormValue("name") != "" {
		greet = greeter.Greet(r.FormValue("name"))
	}

	// prepare the data
	data := struct {
		Greet string
	}{
		Greet: greet,
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
		calculation = calculator.Calculate(r.FormValue("calculation"))
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
	http.Handle("/api/greet", &apiHandler{})
	http.HandleFunc("/greet", htmlHandler)
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
	var addr = flag.String("addr", ":8080", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	startServer(*addr)
}
