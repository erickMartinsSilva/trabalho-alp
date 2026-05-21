package main

import (
	"loadbalancer/src/models"
	"net/http"
)


func (server *models.RoundRobin) IsAlive *models.Backend {
		for _, backend := range server.Backends {
			backend.Mux.RLock()
			err := http.Get(backend.URL.String() + "/health")
			backend.Alive = (err == nil)
		}
}


func rotate(server *models.RoundRobin) *models.Backend {
	server.Mux.Lock()
	server.Current = (server.Current + 1) % len(server.Backends)
	server.Mux.Unlock()

	return server.Backends[server.Current]
}

func getNextAvailableServer(server *models.RoundRobin) *models.Backend {
	for i := 0; i < len(server.Backends); i++ {
		backend := rotate(server)
		if server.backend.IsAlive() {
			return backend
		}
	}
	return nil
}