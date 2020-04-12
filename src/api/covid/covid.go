package covid

import(
	"errors"
	"fmt"
	)

type DataSet struct{
	Data []DataSetRow
	Index string
}

func (d DataSet) Valid() error{
	if d.Index == ""{
		return errors.New("index is nil")
	}
	if len(d.Data) == 0{
		return errors.New("no data in dataset")
	}
	return nil
}

type DataSetRow struct{
	Province string
	Country string
	Latitude string
	Longitud string
	Values []PerDayValues
}

type PerDayValues struct{
	Date string
	Count int
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
