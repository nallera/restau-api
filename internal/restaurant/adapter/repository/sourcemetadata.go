package repository

import (
	"restauAPI/server"
	"time"
)

const RefreshTime = 6 * time.Hour

type WebSourceMetadata interface {
	IsTooOld() bool
	EtagEquals(etag string) bool
	Update(etag string, lastModified time.Time)
}

type webSourceMetadata struct {
	etag         string
	lastModified time.Time
	clock        server.Clock
}

func NewWebSourceMetadata(clock server.Clock) WebSourceMetadata {
	return &webSourceMetadata{
		etag:         "",
		lastModified: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		clock:        clock,
	}
}

func (m *webSourceMetadata) IsTooOld() bool {
	return m.clock.Time().Sub(m.lastModified) > RefreshTime
}

func (m *webSourceMetadata) EtagEquals(etag string) bool {
	return m.etag == etag
}

func (m *webSourceMetadata) Update(etag string, lastModified time.Time) {
	m.etag = etag
	m.lastModified = lastModified
}
