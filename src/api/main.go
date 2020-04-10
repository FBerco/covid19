package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/FBerco/covid19/src/api/server"

)

func main(){
	router := mux.NewRouter()

	server.AppendControllers(router)

	http.ListenAndServe(":80", router)
	http.Handle("/", router)
}
