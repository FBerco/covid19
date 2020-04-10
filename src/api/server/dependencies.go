package server

import "github.com/mercadolibre/covid19/src/api/covid"

func newCovidController() *covid.Controller{
	return &covid.Controller{}
}
