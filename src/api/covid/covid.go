package covid

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