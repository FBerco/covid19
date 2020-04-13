package covid

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type dataService interface {
	GetConfirmed() (DataSet, error)
	GetDeath() (DataSet, error)
	GetRecovered() (DataSet, error)
}

type elasticService interface{
	IndexDataSet(dataset DataSet) error
}

type Controller struct{
	DataService dataService
	ElasticService elasticService
}

func (c Controller) UpdateCases(writer http.ResponseWriter, request *http.Request){
	var dataset DataSet
	var err error
	vars := mux.Vars(request)
	switch
	{
		case vars["case"] == "confirmed":
			dataset, err = c.DataService.GetConfirmed()
			break
		case vars["case"] == "deaths":
			dataset, err = c.DataService.GetDeath()
			break
		case vars["case"] == "recovered":
			dataset, err = c.DataService.GetRecovered()
			break
		default:
			writer.WriteHeader(http.StatusNotFound)
	}
	if err != nil{
		log.Fatalf("Error getting response: %s", err)
	}
	c.ElasticService.IndexDataSet(dataset)
	//fmt.Fprintln(writer, dataset)
}

func (c Controller) RunElastic(writer http.ResponseWriter, request *http.Request){
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fmt.Fprintln(writer, res)
}