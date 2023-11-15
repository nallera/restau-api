package restaurant

import (
	"restauAPI/server"
	"time"
)

const memCacheTTL = 5 * time.Minute

type RestaurantCache interface {
	Get() []*Restaurant
	Update(restaurants []*Restaurant)
	UpdateLastModified()
	DataIsOld() bool
}

type restaurantCache struct {
	restaurants  []*Restaurant
	lastModified time.Time
	clock        server.Clock
}

func NewRestaurantCache(clock server.Clock) RestaurantCache {
	return &restaurantCache{
		restaurants:  nil,
		lastModified: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		clock:        clock,
	}
}

func (rc *restaurantCache) Get() []*Restaurant {
	return rc.restaurants
}

func (rc *restaurantCache) Update(restaurants []*Restaurant) {
	rc.restaurants = restaurants
	rc.lastModified = rc.clock.Time()
}

func (rc *restaurantCache) UpdateLastModified() {
	rc.lastModified = rc.clock.Time()
}

func (rc *restaurantCache) DataIsOld() bool {
	return rc.clock.Time().Sub(rc.lastModified) > memCacheTTL
}
