package covid

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"log"
	"net/http"
)

type Controller struct{

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

	fmt.Println(res)
}