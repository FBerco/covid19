package server
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const(
	GET = "GET"
	POST = "POST"
)

func AppendControllers(router *mux.Router){
	covidController := newCovidController()

	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "up and going")
	}).Methods(GET)

	router.HandleFunc("/elastic", covidController.RunElastic).Methods(GET)
	router.HandleFunc("/update/{case}", covidController.UpdateCases).Methods(GET)
}