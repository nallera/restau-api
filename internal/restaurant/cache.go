package restaurant

import (
	"fmt"
	"restauAPI/server"
	"time"
)

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
	ttl          time.Duration
}

func NewRestaurantCache(clock server.Clock, TTLSeconds int) RestaurantCache {
	ttl := time.Duration(TTLSeconds) * time.Second
	println(fmt.Sprintf("Restaurant cache TTL: %+v", ttl))
	return &restaurantCache{
		restaurants:  nil,
		lastModified: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		clock:        clock,
		ttl:          ttl,
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
	return rc.clock.Time().Sub(rc.lastModified) > rc.ttl
}
