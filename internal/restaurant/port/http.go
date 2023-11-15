package port

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restauAPI/internal/restaurant/app"
	"restauAPI/server"
)

type Handler interface {
	GetAvailableRestaurants(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	restaurantsService app.RestaurantService
}

func NewHTTPHandler(restaurantsService app.RestaurantService) Handler {
	return &handler{
		restaurantsService: restaurantsService,
	}
}

type LocationMessage struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (h *handler) GetAvailableRestaurants(w http.ResponseWriter, r *http.Request) {
	locationMessage := new(LocationMessage)

	err := server.ExtractMessage(r, locationMessage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "error parsing message: %v", err)
		return
	}

	restaurantIDs, err := h.restaurantsService.GetAvailableRestaurants(locationMessage.Latitude, locationMessage.Longitude)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "failed to get available restaurants: %v", err)
		return
	}

	json.NewEncoder(w).Encode(restaurantIDs)
}
