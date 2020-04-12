package server

import (
	"github.com/FBerco/covid19/src/api/covid"
	"log"
	"github.com/elastic/go-elasticsearch/v8"

)

func newCovidController() *covid.Controller{
	return &covid.Controller{
		DataService: newDataService(),
		ElasticService: newElasticService(),
	}
}

func newDataService() *covid.DataService{
	return &covid.DataService{}
}

func newElasticService() *covid.ElasticService{
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return &covid.ElasticService{
		Client: *es,
	}
}