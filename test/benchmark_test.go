package test

import (
	"fmt"
	"restauAPI/internal/restaurant/app"
	"testing"
	"time"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
	{input: 1000000},
	{input: 10000000},
	{input: 100000000},
}

func BenchmarkRestaurantServiceFilterRestaurants(b *testing.B) {
	latitude := 50.12
	longitude := 60.87
	testTime := time.Date(2023, 11, 13, 20, 25, 0, 0, time.UTC)

	for _, v := range table {
		restaurantsResponse := MakeRestaurantsForBenchmark(v.input)
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				app.FilterRestaurants(restaurantsResponse, latitude, longitude, testTime)
			}
		})
	}
}
