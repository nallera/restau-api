package repository

//
//import (
//	"encoding/json"
//	"fmt"
//	"restauAPI/internal/restaurant"
//	"restauAPI/server"
//)
//
//type CacheItem struct {
//	Restaurants []*restaurant.Restaurant `json:"restaurants"`
//	Etag        string                   `json:"etag"`
//}
//
//type CacheClient struct {
//	cacheClient server.CacheClient
//	restClient RestRepositoryClient
//}
//
//func NewCacheClient(cacheClient server.CacheClient) *CacheClient {
//	return &CacheClient{
//		cacheClient: cacheClient,
//	}
//}
//
//func (r *CacheClient) GetRestaurants() ([]*restaurant.Restaurant, error) {
//	data, getErr := r.cacheClient.Get()
//	var (
//		cacheItem CacheItem
//		oldEtag   string
//	)
//
//	if data != "" {
//		err := json.Unmarshal([]byte(data), &cacheItem)
//		if err != nil {
//			return []*restaurant.Restaurant{}, err
//		}
//		oldEtag = cacheItem.Etag
//	}
//
//	switch getErr.(type) {
//	case *server.ErrCacheNotFound:
//		restaurants, etag, restErr := r.restClient.GetRestaurants()
//		if restErr != nil {
//			return nil, fmt.Errorf("failed to get restaurants: %v", restErr)
//		}
//
//		if oldEtag != etag {
//			r.storeInCache(restaurants, etag)
//		}
//
//		return restaurants, nil
//
//	case *server.ErrCacheUnknown:
//		return nil, fmt.Errorf("failed to get restaurants: %v", getErr)
//	}
//
//	return cacheItem.Restaurants, nil
//}
//
//func (r *CacheClient) storeInCache(restaurants []*restaurant.Restaurant, etag string) {
//	cacheItem := CacheItem{
//		Restaurants: restaurants,
//		Etag:        etag,
//	}
//	jsonCacheItem, _ := json.Marshal(cacheItem)
//
//	_, cacheErr := r.cacheClient.Store(string(jsonCacheItem))
//	if cacheErr != nil {
//		println(fmt.Errorf("failed to store in cache: %v", cacheErr))
//	}
//}
