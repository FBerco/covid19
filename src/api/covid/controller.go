package covid

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

type DataService interface {
	GetConfirmedCases() ([]DataSetRow, error)
	GetDeath() ([]DataSetRow, error)
	GetRecovered() ([]DataSetRow, error)
}

type Controller struct{
	Service DataService
}

func (c Controller) DownloadData(writer http.ResponseWriter, request *http.Request){
	dataset, err := c.Service.GetConfirmedCases()
	if err != nil{
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Fprintln(writer, dataset)
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