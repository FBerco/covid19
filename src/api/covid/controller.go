package covid

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

type dataService interface {
	GetConfirmedCases() (DataSet, error)
	GetDeath() (DataSet, error)
	GetRecovered() (DataSet, error)
}

type elasticService interface{
	IndexDataSet(dataset DataSet)
}

type Controller struct{
	DataService dataService
	ElasticService elasticService
}

func (c Controller) DownloadData(writer http.ResponseWriter, request *http.Request){
	dataset, err := c.DataService.GetConfirmedCases()
	if err != nil{
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println("Got the data")
	fmt.Println("Indexing to elastic")
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