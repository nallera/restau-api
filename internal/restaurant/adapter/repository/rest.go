package repository

import (
	"fmt"
	"restauAPI/internal/restaurant"
	"restauAPI/server"
)

const (
	RestaurantsBaseURL = "https://s3.amazonaws.com/test.jampp.com/dmarasca"
)

type RestRepositoryClient struct {
	restClient        server.RestClient
	webSourceMetadata WebSourceMetadata
}

func NewRestRepositoryClient(restClient server.RestClient, webSourceMetadata WebSourceMetadata) *RestRepositoryClient {
	return &RestRepositoryClient{
		restClient:        restClient,
		webSourceMetadata: webSourceMetadata,
	}
}

func (r *RestRepositoryClient) GetRestaurants() ([]*restaurant.Restaurant, error) {
	if r.webSourceMetadata.IsTooOld() { // if the web source metadata is too old according to the estimated update frequency
		url := fmt.Sprintf("%s/takehome.csv", RestaurantsBaseURL)

		headers, err := r.restClient.GetHead(url) // get the headers to check for changes before downloading the data
		if err != nil {
			return nil, fmt.Errorf("failed to get headers for restaurants: %v", err)
		}

		// if no change in the web source data, return nil
		if r.webSourceMetadata.EtagEquals(headers.Etag) {
			return nil, nil
		}

		// if the data changed, update the metadata
		r.webSourceMetadata.Update(headers.Etag, headers.LastModified)

		csvLines, err := r.restClient.GetCSV(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get restaurants: %v", err)
		}

		restaurants := restaurant.CSVToAppRestaurants(csvLines)

		return restaurants, nil
	}

	return nil, nil
}
