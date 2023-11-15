package restaurant

import (
	"strconv"
	"time"
)

type Restaurant struct {
	ID                 uint64
	Latitude           float64
	Longitude          float64
	AvailabilityRadius float64
	OpenHour           time.Time
	CloseHour          time.Time
}

func CSVToAppRestaurant(csv []string) *Restaurant {
	id, _ := strconv.ParseUint(csv[0], 10, 64)
	latitude, _ := strconv.ParseFloat(csv[1], 64)
	longitude, _ := strconv.ParseFloat(csv[2], 64)
	availabilityRadius, _ := strconv.ParseFloat(csv[3], 64)
	openHour, _ := time.Parse(time.TimeOnly, csv[4])
	closeHour, _ := time.Parse(time.TimeOnly, csv[5])

	return &Restaurant{
		ID:                 id,
		Latitude:           latitude,
		Longitude:          longitude,
		AvailabilityRadius: availabilityRadius,
		OpenHour:           openHour,
		CloseHour:          closeHour,
	}
}

func CSVToAppRestaurants(csvLines [][]string) []*Restaurant {
	var restaurants []*Restaurant

	for _, csvLine := range csvLines {
		restaurants = append(restaurants, CSVToAppRestaurant(csvLine))
	}

	return restaurants
}

func AppRestaurantToCSV(restaurant *Restaurant) []string {
	var csvString []string

	csvString = append(csvString, strconv.FormatUint(restaurant.ID, 10))
	csvString = append(csvString, strconv.FormatFloat(restaurant.Latitude, 10, 8, 64))
	csvString = append(csvString, strconv.FormatFloat(restaurant.Longitude, 10, 8, 64))
	csvString = append(csvString, strconv.FormatFloat(restaurant.AvailabilityRadius, 10, 2, 64))
	csvString = append(csvString, restaurant.OpenHour.String())
	csvString = append(csvString, restaurant.CloseHour.String())

	return csvString
}

func AppRestaurantsToCSV(restaurants []*Restaurant) [][]string {
	var csvStrings [][]string

	for _, r := range restaurants {
		csvStrings = append(csvStrings, AppRestaurantToCSV(r))
	}

	return csvStrings
}
