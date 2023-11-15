package test

import (
	"restauAPI/internal/restaurant"
	"time"
)

func MakeRestaurantsResponse() []*restaurant.Restaurant {
	return []*restaurant.Restaurant{
		{
			ID:                 1,
			Latitude:           50,
			Longitude:          60,
			AvailabilityRadius: 4,
			OpenHour:           time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC),
			CloseHour:          time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC),
		},
		{
			ID:                 2,
			Latitude:           55,
			Longitude:          70,
			AvailabilityRadius: 15,
			OpenHour:           time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC),
			CloseHour:          time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC),
		},
		{
			ID:                 3,
			Latitude:           55,
			Longitude:          70,
			AvailabilityRadius: 2,
			OpenHour:           time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC),
			CloseHour:          time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC),
		},
		{
			ID:                 4,
			Latitude:           50,
			Longitude:          60,
			AvailabilityRadius: 5,
			OpenHour:           time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC),
			CloseHour:          time.Date(0, 1, 1, 10, 0, 0, 0, time.UTC),
		},
	}
}
