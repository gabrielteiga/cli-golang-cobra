package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielteiga/golang-cli-cobra/app"
)

type Server struct {
	Port string
}

func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		a = 0
	}

	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}

	calc := app.NewCalc()
	calc.B = b
	calc.A = a

	w.Write([]byte(fmt.Sprintf("O resultado é: %f", calc.Sum())))
}

func (s Server) Serve() {
	http.HandleFunc("/", s.SumHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
