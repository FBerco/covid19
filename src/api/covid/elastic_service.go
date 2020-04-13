package covid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type ElasticService struct{
	Client elasticsearch.Client
}

func (service ElasticService) IndexDataSet(dataset DataSet) error{
	if err := dataset.Valid(); err != nil{
		return err
	}
	ctx := context.Background()

	for _, data := range dataset.Data{
		doc := ElasticDoc{
			Province: data.Province,
			Country: data.Country,
			Latitude: data.Latitude,
			Longitud: data.Longitud,
		}
		for  _, perday := range data.Values{
			doc.Date = perday.Date
			doc.Count = perday.Count

			//todo GO ROUTINE
			service.IndexDoc(ctx, doc, dataset.Index)
		}
	}
	return nil
}

func (service ElasticService) IndexDoc(ctx context.Context, doc ElasticDoc, index string) {
	body, _ := json.Marshal(doc)
	req := esapi.IndexRequest{
		Index: index,
		DocumentID: doc.GetId(),
		Body: strings.NewReader(string(body)),
		Refresh: "true",
	}

	res, err := req.Do(ctx, service.Client.Transport)
	if err != nil {
		log.Fatalf("IndexRequest ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("%s ERROR indexing document ID=%s", res.Status(), req.DocumentID)
	} else {
		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("\nIndexRequest() RESPONSE:")
			fmt.Println("Status:", res.Status())
			fmt.Println("Result:", resMap["result"])
			fmt.Println("Version:", int(resMap["_version"].(float64)))
			fmt.Println("resMap:", resMap)
		}
	}
}
