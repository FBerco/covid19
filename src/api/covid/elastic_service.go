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

const (
	index = "some_index"
)
type ElasticService struct{
	Client elasticsearch.Client
}

type ElasticDoc struct{
	Province string
	Country string
	Latitude string
	Longitud string
	Date string
	Count int
}

func (e ElasticDoc) GetId() string{
	if e.Province != ""{
		return fmt.Sprintf("%s-%s-%s", e.Date, e.Country, e.Province )
	}
	return fmt.Sprintf("%s-%s", e.Date, e.Country)
}

func (service ElasticService) IndexDataSet(dataset []DataSetRow){
	ctx := context.Background()

	for _, data := range dataset{
		doc := ElasticDoc{
			Province: data.Province,
			Country: data.Country,
			Latitude: data.Latitude,
			Longitud: data.Longitud,
		}
		for  _, perday := range data.Values{
			doc.Date = perday.Date
			doc.Count = perday.Count

			go service.IndexDoc(ctx, doc)
		}
	}
}

func (service ElasticService) IndexDoc(ctx context.Context, doc ElasticDoc) {
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
		log.Printf("%s ERROR indexing document ID=%d", res.Status(), req.DocumentID)
	} else {
		// Deserialize the response into a map.
		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("\nIndexRequest() RESPONSE:")
			// Print the response status and indexed document version.
			fmt.Println("Status:", res.Status())
			fmt.Println("Result:", resMap["result"])
			fmt.Println("Version:", int(resMap["_version"].(float64)))
			fmt.Println("resMap:", resMap)
			fmt.Println("\n")
		}
	}
}

