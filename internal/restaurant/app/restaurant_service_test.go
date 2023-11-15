package app_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"restauAPI/internal/restaurant/app"
	"restauAPI/test"
	"testing"
	"time"
)

func TestRestaurantService_GetAvailableRestaurants(t *testing.T) {
	type depFields struct {
		restaurantRepoMock  *test.RestaurantRepositoryMock
		restaurantCacheMock *test.RestaurantCacheMock
		clockMock           *test.ClockMock
	}
	type input struct {
		latitude  float64
		longitude float64
	}
	type output struct {
		restaurantIDs []uint64
		err           error
	}

	latitude := 50.12
	longitude := 60.87

	restaurantsResponse := test.MakeRestaurantsResponse()
	filteredRestaurants := []uint64{1, 2}
	testTime := time.Date(2023, 11, 13, 20, 25, 0, 0, time.UTC)

	tests := []struct {
		name   string
		in     input
		on     func(*depFields)
		assert func(*testing.T, *output)
	}{
		{
			name: "return only available restaurants",
			in: input{
				latitude:  latitude,
				longitude: longitude,
			},
			on: func(df *depFields) {
				df.restaurantCacheMock.On("DataIsOld").Return(true).Once()
				df.restaurantRepoMock.On("GetRestaurants").Return(restaurantsResponse, nil).Once()
				df.restaurantCacheMock.On("Update").Once()
				df.restaurantCacheMock.On("Get").Return(restaurantsResponse).Once()
				df.clockMock.On("Time").Return(testTime, nil).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.NoError(t, out.err)
				assert.Equal(t, filteredRestaurants, out.restaurantIDs)
			},
		},
		{
			name: "repo error",
			in: input{
				latitude:  latitude,
				longitude: longitude,
			},
			on: func(df *depFields) {
				df.restaurantCacheMock.On("DataIsOld").Return(true).Once()
				df.restaurantRepoMock.On("GetRestaurants").Return(nil, errors.New("repo error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "error getting restaurants for location")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Having
			restaurantRepoMock := new(test.RestaurantRepositoryMock)
			clockMock := new(test.ClockMock)
			restaurantCacheMock := new(test.RestaurantCacheMock)
			s := app.NewRestaurantService(restaurantRepoMock, restaurantCacheMock, clockMock)

			f := &depFields{restaurantRepoMock: restaurantRepoMock, restaurantCacheMock: restaurantCacheMock, clockMock: clockMock}

			tt.on(f)

			// When
			restaurantIDs, err := s.GetAvailableRestaurants(tt.in.latitude, tt.in.longitude)

			o := output{restaurantIDs: restaurantIDs, err: err}
			// Then
			tt.assert(t, &o)
			restaurantRepoMock.AssertExpectations(t)
			restaurantCacheMock.AssertExpectations(t)
			clockMock.AssertExpectations(t)
		})
	}
}
