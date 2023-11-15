package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"restauAPI/internal/restaurant"
	"restauAPI/internal/restaurant/adapter/repository"
	"restauAPI/server"
	"restauAPI/test"
	"testing"
	"time"
)

func TestRepoRestClient_GetRestaurants(t *testing.T) {
	type depFields struct {
		restClient            *test.RestClientMock
		webSourceMetadataMock *test.WebSourceMetadataMock
	}

	type output struct {
		restaurants []*restaurant.Restaurant
		err         error
	}

	restaurantsResponse := restaurant.AppRestaurantsToCSV(test.MakeRestaurantsResponse())
	url := repository.RestaurantsBaseURL + "/takehome.csv"
	testTime := time.Date(2023, 11, 13, 20, 25, 0, 0, time.UTC)
	testHeaders := server.Headers{
		Etag:         "test-etag",
		LastModified: testTime,
	}

	tests := []struct {
		name   string
		on     func(*depFields)
		assert func(*testing.T, *output)
	}{
		{
			name: "get successful",
			on: func(df *depFields) {
				df.webSourceMetadataMock.On("IsTooOld").Return(true).Once()
				df.restClient.On("GetHead", url).Return(testHeaders, nil).Once()
				df.webSourceMetadataMock.On("EtagEquals").Return(false).Once()
				df.restClient.On("GetCSV", url).Return(restaurantsResponse, nil).Once()
				df.webSourceMetadataMock.On("Update").Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.NoError(t, out.err)
			},
		},
		{
			name: "get head failed",
			on: func(df *depFields) {
				df.webSourceMetadataMock.On("IsTooOld").Return(true).Once()
				df.restClient.On("GetHead", url).Return(testHeaders, errors.New("test error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "failed to get headers for restaurants: test error")
			},
		},
		{
			name: "same etag",
			on: func(df *depFields) {
				df.webSourceMetadataMock.On("IsTooOld").Return(true).Once()
				df.restClient.On("GetHead", url).Return(testHeaders, nil).Once()
				df.webSourceMetadataMock.On("EtagEquals").Return(true).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.NoError(t, out.err)
			},
		},
		{
			name: "get head successful error getting csv",
			on: func(df *depFields) {
				df.webSourceMetadataMock.On("IsTooOld").Return(true).Once()
				df.restClient.On("GetHead", url).Return(testHeaders, nil).Once()
				df.webSourceMetadataMock.On("EtagEquals").Return(false).Once()
				df.webSourceMetadataMock.On("Update").Once()
				df.restClient.On("GetCSV", url).Return(nil, errors.New("test error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "failed to get restaurants: test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Having
			restClient := new(test.RestClientMock)
			webSourceMetadataMock := new(test.WebSourceMetadataMock)
			restRepository := repository.NewRestRepositoryClient(restClient, webSourceMetadataMock)

			f := &depFields{restClient: restClient, webSourceMetadataMock: webSourceMetadataMock}

			tt.on(f)

			// When
			_, err := restRepository.GetRestaurants()

			// Then
			tt.assert(t, &output{err: err})
			restClient.AssertExpectations(t)
			webSourceMetadataMock.AssertExpectations(t)
		})
	}
}
