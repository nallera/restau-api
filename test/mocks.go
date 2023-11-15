package test

import (
	"github.com/stretchr/testify/mock"
	"restauAPI/internal/restaurant"
	"restauAPI/server"
	"time"
)

type RestClientMock struct {
	mock.Mock
}

func (m *RestClientMock) GetCSV(url string) ([][]string, error) {
	args := m.Called(url)
	d, _ := args.Get(0).([][]string)
	err, _ := args.Get(1).(error)

	return d, err
}

func (m *RestClientMock) GetHead(url string) (server.Headers, error) {
	args := m.Called(url)
	d, _ := args.Get(0).(server.Headers)
	err, _ := args.Get(1).(error)

	return d, err
}

type CacheClientMock struct {
	mock.Mock
}

func (m *CacheClientMock) Get() (string, error) {
	args := m.Called()
	d, _ := args.Get(0).(string)
	e, _ := args.Get(1).(error)

	return d, e
}

func (m *CacheClientMock) Store(_ string) (string, error) {
	args := m.Called()
	d, _ := args.Get(0).(string)
	e, _ := args.Get(1).(error)

	return d, e
}

type RestaurantRepositoryMock struct {
	mock.Mock
}

func (m *RestaurantRepositoryMock) GetRestaurants() ([]*restaurant.Restaurant, error) {
	args := m.Called()
	d, _ := args.Get(0).([]*restaurant.Restaurant)
	e, _ := args.Get(1).(error)

	return d, e
}

type ClockMock struct {
	mock.Mock
}

func (m *ClockMock) Time() time.Time {
	args := m.Called()
	t, _ := args.Get(0).(time.Time)

	return t
}

type RestaurantCacheMock struct {
	mock.Mock
}

func (m *RestaurantCacheMock) Get() []*restaurant.Restaurant {
	args := m.Called()
	d, _ := args.Get(0).([]*restaurant.Restaurant)

	return d
}

func (m *RestaurantCacheMock) Update(_ []*restaurant.Restaurant) {
	_ = m.Called()

	return
}

func (m *RestaurantCacheMock) UpdateLastModified() {
	_ = m.Called()

	return
}

func (m *RestaurantCacheMock) DataIsOld() bool {
	args := m.Called()
	b, _ := args.Get(0).(bool)

	return b
}

type WebSourceMetadataMock struct {
	mock.Mock
}

func (m *WebSourceMetadataMock) IsTooOld() bool {
	args := m.Called()
	b, _ := args.Get(0).(bool)

	return b
}

func (m *WebSourceMetadataMock) EtagEquals(etag string) bool {
	args := m.Called()
	b, _ := args.Get(0).(bool)

	return b
}

func (m *WebSourceMetadataMock) Update(etag string, lastModified time.Time) {
	_ = m.Called()

	return
}
