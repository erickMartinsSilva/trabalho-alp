package models

import (
	"sync"
)

type RoundRobin struct{
	Backends []*Backend
	Current int
	Mux sync.Mutex
}