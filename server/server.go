package server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

const (
	ListKey  = "listkey"
	RedisTTL = 10 * time.Minute
)

type Headers struct {
	Etag         string
	LastModified time.Time
}

type RestClient interface {
	GetCSV(url string) ([][]string, error)
	GetHead(url string) (Headers, error)
}

func NewRestClient() RestClient {
	return &restClient{}
}

type restClient struct{}

func (r *restClient) GetCSV(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d (%s)", resp.StatusCode, resp.Body)
	}

	csvReader := csv.NewReader(resp.Body)

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV lines: %w", err)
	}

	// remove header
	lines = lines[1:]

	return lines, nil
}

func (r *restClient) GetHead(url string) (Headers, error) {
	headers := Headers{}

	resp, err := http.Head(url)
	if err != nil {
		return headers, fmt.Errorf("failed to make GET HEAD request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return headers, fmt.Errorf("unexpected status code: %d (%s)", resp.StatusCode, resp.Body)
	}

	headers.Etag = resp.Header["Etag"][0]
	headers.LastModified, err = time.Parse(time.RFC1123, resp.Header["Last-Modified"][0])
	if err != nil {
		return headers, fmt.Errorf("failed to parse GET HEAD request: %w", err)
	}

	return headers, nil
}

type CacheClient interface {
	Get() (string, error)
	Store(data string) (string, error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) CacheClient {
	return &redisClient{
		client: client,
	}
}

func (r *redisClient) Get() (string, error) {
	data, err := r.client.Get(context.Background(), ListKey).Result()
	if err == redis.Nil {
		return data, NewErrCacheNotFound(fmt.Sprintf("not found for key %s", ListKey))
	} else if err != nil {
		return data, NewErrCacheUnknown(fmt.Sprintf("get error for key %s: %v", ListKey, err))
	}

	return data, nil
}

func (r *redisClient) Store(data string) (string, error) {
	data, err := r.client.Set(context.Background(), ListKey, data, RedisTTL).Result()

	if err != nil {
		return data, NewErrCacheUnknown(fmt.Sprintf("store error for key %s: %v", ListKey, err))
	}

	return data, nil
}

func ExtractMessage(r *http.Request, msg interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(msg); err != nil {
		return errors.New(fmt.Sprintf("error reading request body: %s", err))
	}

	return nil
}
