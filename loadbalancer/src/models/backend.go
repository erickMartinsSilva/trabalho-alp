package models

import (
	"net/url"
	"sync"
)

type Backend struct {
	URL	*url.URL
	Alive bool
	mux sync.RWMutex
}