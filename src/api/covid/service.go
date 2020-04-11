package covid

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	confirmedUrl = "https://data.humdata.org/hxlproxy/api/data-preview.csv?url=https%3A%2F%2Fraw.githubusercontent.com%2FCSSEGISandData%2FCOVID-19%2Fmaster%2Fcsse_covid_19_data%2Fcsse_covid_19_time_series%2Ftime_series_covid19_confirmed_global.csv&filename=time_series_covid19_confirmed_global.csv"
	deathsUrl = "https://data.humdata.org/hxlproxy/api/data-preview.csv?url=https%3A%2F%2Fraw.githubusercontent.com%2FCSSEGISandData%2FCOVID-19%2Fmaster%2Fcsse_covid_19_data%2Fcsse_covid_19_time_series%2Ftime_series_covid19_deaths_global.csv&filename=time_series_covid19_deaths_global.csv"
	recoverdUrl = "https://data.humdata.org/hxlproxy/api/data-preview.csv?url=https%3A%2F%2Fraw.githubusercontent.com%2FCSSEGISandData%2FCOVID-19%2Fmaster%2Fcsse_covid_19_data%2Fcsse_covid_19_time_series%2Ftime_series_covid19_recovered_global.csv&filename=time_series_covid19_recovered_global.csv"
	)


type Service struct{

}

func (s Service) GetConfirmedCases() ([]DataSetRow, error){
	return s.getData(confirmedUrl)
}
func (s Service) GetDeath() ([]DataSetRow, error){
	return s.getData(deathsUrl)
}
func (s Service) GetRecovered() ([]DataSetRow, error){
	return s.getData(recoverdUrl)
}

func (s Service) getData(url string) ([]DataSetRow, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading data: %s", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var dataset []DataSetRow
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan(){
		data := strings.Split(scanner.Text(), ",")
		row := DataSetRow{
			Province: data[0],
			Country: data[1],
			Latitude: data[2],
			Longitud: data[3],
		}
		date := time.Date(2020, 01, 22,0,0,0,0, time.UTC)
		var values []PerDayValues
		for i := 5; i < len(data); i++{
			value, err := strconv.Atoi(data[i])
			if err != nil {
				fmt.Println(data)
				return nil, fmt.Errorf("error converting data: %s, i: %s", err, i)
				}
			values = append(values, PerDayValues{
				Date:  date.Format("2006-01-02"),
				Count: value,
			})
			date.AddDate(0,0,1)
		}
		row.Values = values
		dataset = append(dataset, row)
	}
	return dataset, nil
}