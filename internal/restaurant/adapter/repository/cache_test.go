package repository_test

//
//import (
//	"encoding/json"
//	"github.com/stretchr/testify/assert"
//	"restauAPI/internal/restaurant"
//	"restauAPI/internal/restaurant/adapter/repository"
//	"restauAPI/server"
//	"restauAPI/test"
//	"testing"
//)
//
//func TestCacheClient_GetRestaurants(t *testing.T) {
//	type output struct {
//		restaurants []*restaurant.Restaurant
//		etag        string
//		err         error
//	}
//	restaurantsResponse := restaurant.AppRestaurantsToCSV(test.MakeRestaurantsResponse())
//	url := repository.RestaurantsBaseURL + "/takehome.csv"
//	etag := "test-etag"
//
//	cacheItem := repository.CacheItem{
//		Restaurants: test.MakeRestaurantsResponse(),
//		Etag:        etag,
//	}
//	jsonCacheItem, _ := json.Marshal(cacheItem)
//
//	tests := []struct {
//		name   string
//		on     func(*depFieldsCache)
//		assert func(*testing.T, *output)
//	}{
//		{
//			name: "get successful from web (empty cache)",
//			on: func(df *depFieldsCache) {
//				df.cacheClient.On("Get").Return("", server.NewErrCacheNotFound("test message")).Once()
//				df.restClient.On("GetCSV", url).Return(restaurantsResponse, etag, nil).Once()
//				df.cacheClient.On("Store").Return("", nil).Once()
//			},
//			assert: func(t *testing.T, out *output) {
//				assert.NoError(t, out.err)
//			},
//		},
//		{
//			name: "get successful from web (out of date cache)",
//			on: func(df *depFieldsCache) {
//				df.cacheClient.On("Get").Return(string(jsonCacheItem), server.NewErrCacheNotFound("test message")).Once()
//				df.restClient.On("GetCSV", url).Return(restaurantsResponse, etag, nil).Once()
//				df.cacheClient.On("Store").Return("", nil).Once()
//			},
//			assert: func(t *testing.T, out *output) {
//				assert.NoError(t, out.err)
//			},
//		},
//		{
//			name: "get successful from cache",
//			on: func(df *depFieldsCache) {
//				df.cacheClient.On("Get").Return(string(jsonCacheItem), nil).Once()
//			},
//			assert: func(t *testing.T, out *output) {
//				assert.NoError(t, out.err)
//			},
//		},
//		{
//			name: "get failed from cache",
//			on: func(df *depFieldsCache) {
//				df.cacheClient.On("Get").Return("", server.NewErrCacheUnknown("test message")).Once()
//			},
//			assert: func(t *testing.T, out *output) {
//				assert.ErrorContains(t, out.err, "failed to get restaurants: test message")
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Having
//			cacheClient := new(test.CacheClientMock)
//			restClient := new(test.RestClientMock)
//			repoRestClient := repository.NewRestRepositoryClient(restClient)
//			restRepository := repository.NewCacheClient(cacheClient)
//
//			f := &depFieldsCache{cacheClient: cacheClient, restClient: restClient}
//
//			tt.on(f)
//
//			// When
//			_, err := restRepository.GetRestaurants()
//
//			// Then
//			tt.assert(t, &output{err: err})
//			restClient.AssertExpectations(t)
//		})
//	}
//}
//
//type depFieldsCache struct {
//	cacheClient *test.CacheClientMock
//	restClient  *test.RestClientMock
//}
