package app

import (
	"fmt"
	"math"
	"restauAPI/internal/restaurant"
	"restauAPI/server"
	"time"
)

type RestaurantService interface {
	GetAvailableRestaurants(latitude, longitude float64) ([]uint64, error)
}

type restaurantService struct {
	restaurantRepository restaurant.Repository
	clock                server.Clock
	restaurantCache      restaurant.RestaurantCache
}

func NewRestaurantService(restaurantRepository restaurant.Repository, restaurantCache restaurant.RestaurantCache, clock server.Clock) RestaurantService {
	return &restaurantService{
		restaurantRepository: restaurantRepository,
		clock:                clock,
		restaurantCache:      restaurantCache,
	}
}

func (rs *restaurantService) GetAvailableRestaurants(latitude, longitude float64) ([]uint64, error) {
	err := rs.updateCacheIfNeeded()
	if err != nil {
		return nil, fmt.Errorf("error getting restaurants for location (%.5f, %.5f): %v", latitude, longitude, err)
	}

	// filter restaurants by the business needs
	availableRestaurants := FilterRestaurants(rs.restaurantCache.Get(), latitude, longitude, getCurrentTimeNoDate(rs.clock.Time()))

	return availableRestaurants, nil
}

func (rs *restaurantService) updateCacheIfNeeded() error {
	if rs.restaurantCache.DataIsOld() { // if the data in the cache is older than the accepted margin
		restaurants, repoErr := rs.restaurantRepository.GetRestaurants() // fetch from web
		if repoErr != nil {
			return repoErr
		}
		if restaurants != nil { // if there's new data, update the cache and the last modified date
			rs.restaurantCache.Update(restaurants)
		} else { // if there's no new data, update the last modified to reset the TTL
			rs.restaurantCache.UpdateLastModified()
		}
	}
	return nil
}

func FilterRestaurants(restaurants []*restaurant.Restaurant, latitude, longitude float64, currentTime time.Time) []uint64 {
	var aux []*restaurant.Restaurant
	var availableRestaurants []uint64

	for _, r := range restaurants {
		if !EnoughTimeToDeliver(&currentTime, &r.OpenHour, &r.CloseHour) {
			continue
		}
		if !WithinDeliveryRange(latitude, longitude, r.Latitude, r.Longitude, r.AvailabilityRadius) {
			continue
		}

		aux = append(aux, r)
		availableRestaurants = append(availableRestaurants, r.ID)
	}

	return availableRestaurants
}

func getCurrentTimeNoDate(currentTime time.Time) time.Time {
	currentTimeNoDate := time.Date(0, 1, 1, currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, time.UTC)

	return currentTimeNoDate
}

func EnoughTimeToDeliver(currentTime, openHour, closeHour *time.Time) bool {
	return currentTime.After(*openHour) && currentTime.Before(*closeHour)
}

func WithinDeliveryRange(latitude, longitude, rLatitude, rLongitude, radius float64) bool {
	distance := math.Sqrt(math.Pow(latitude-rLatitude, 2) + math.Pow(longitude-rLongitude, 2))

	return distance < radius
}
