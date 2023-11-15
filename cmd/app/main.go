package main

import (
	"fmt"
	"restauAPI/internal/restaurant"

	//"github.com/go-redis/redis/v8"
	"net/http"
	"restauAPI/internal/restaurant/adapter/repository"
	restauapp "restauAPI/internal/restaurant/app"
	restauport "restauAPI/internal/restaurant/port"
	"restauAPI/server"

	"github.com/gorilla/mux"
)

func main() {
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr: "redis-server:6379",
	//	DB:   0,
	//})
	//cacheClient := server.NewRedisClient(redisClient)
	//restaurantRepo := repository.NewCacheClient(cacheClient)

	clock := server.NewClock()
	restClient := server.NewRestClient()
	webSourceMetadata := repository.NewWebSourceMetadata(clock)
	restaurantCache := restaurant.NewRestaurantCache(clock)
	restRepositoryClient := repository.NewRestRepositoryClient(restClient, webSourceMetadata)
	restaurantService := restauapp.NewRestaurantService(restRepositoryClient, restaurantCache, clock)
	restaurantHandler := restauport.NewHTTPHandler(restaurantService)

	r := mux.NewRouter()
	r.HandleFunc("/restaurants/available", restaurantHandler.GetAvailableRestaurants).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
