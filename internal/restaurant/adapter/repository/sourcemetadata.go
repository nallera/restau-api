package repository

import (
	"fmt"
	"restauAPI/server"
	"time"
)

type WebSourceMetadata interface {
	IsTooOld() bool
	EtagEquals(etag string) bool
	Update(etag string, lastModified time.Time)
}

type webSourceMetadata struct {
	etag          string
	lastModified  time.Time
	clock         server.Clock
	refreshPeriod time.Duration
}

func NewWebSourceMetadata(clock server.Clock, refreshPeriodSeconds int) WebSourceMetadata {
	refreshPeriod := time.Duration(refreshPeriodSeconds) * time.Second
	println(fmt.Sprintf("Web Source Metadata refresh period: %+v", refreshPeriod))
	return &webSourceMetadata{
		etag:          "",
		lastModified:  time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		clock:         clock,
		refreshPeriod: refreshPeriod,
	}
}

func (m *webSourceMetadata) IsTooOld() bool {
	return m.clock.Time().Sub(m.lastModified) > m.refreshPeriod
}

func (m *webSourceMetadata) EtagEquals(etag string) bool {
	return m.etag == etag
}

func (m *webSourceMetadata) Update(etag string, lastModified time.Time) {
	m.etag = etag
	m.lastModified = lastModified
}
