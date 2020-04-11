package server

import "github.com/FBerco/covid19/src/api/covid"

func newCovidController() *covid.Controller{
	return &covid.Controller{
		Service: newCovidService(),
	}
}

func newCovidService() *covid.Service{
	return &covid.Service{}
}
