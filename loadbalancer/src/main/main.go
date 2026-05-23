package main

import (
	"fmt"
	"loadbalancer/src/models"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func IsAlive(backend *models.Backend) bool {
	resp, err := http.Get(backend.URL.String() + "/users")
	alive := (err == nil && resp != nil && resp.StatusCode == http.StatusOK)
	if resp != nil {
		resp.Body.Close()
	}

	backend.Mux.Lock()
	backend.Alive = alive
	backend.Mux.Unlock()
	return alive
}

func rotate(server *models.RoundRobin) *models.Backend {
	server.Mux.Lock()
	defer server.Mux.Unlock()

	if len(server.Backends) == 0 {
		return nil
	}
	server.Current = (server.Current + 1) % len(server.Backends)
	return server.Backends[server.Current]
}

func getNextAvailableServer(server *models.RoundRobin) *models.Backend {
	n := len(server.Backends)
	if n == 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		backend := rotate(server)
		if backend == nil {
			continue
		}

		backend.Mux.RLock()
		alive := backend.Alive
		backend.Mux.RUnlock()

		if alive || IsAlive(backend) {
			return backend
		}
	}
	return nil
}

func parseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}

func healthCheck(server *models.RoundRobin) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, backend := range server.Backends {
			go func(b *models.Backend) {
				IsAlive(b)
			}(backend)
		}
	}
}

func main() {

	backends := []*models.Backend{
		{URL: parseURL("http://api1:8080"), Alive: true},
		{URL: parseURL("http://api2:8080"), Alive: true},
		{URL: parseURL("http://api3:8080"), Alive: true},
	}

	lb := &models.RoundRobin{
		Backends: backends,
		Current:  0,
	}


	go healthCheck(lb)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backend := getNextAvailableServer(lb)
		if backend == nil {
			http.Error(w, "Todos os backends estão offline", http.StatusServiceUnavailable)
			return
		}
		fmt.Printf("Roteando %s %s para %s\n", r.Method, r.RequestURI, backend.URL.String())
		reverseProxy := httputil.NewSingleHostReverseProxy(backend.URL)
		reverseProxy.ServeHTTP(w, r)
	})

	fmt.Println("Load Balancer rodando na porta 8080!")
	fmt.Println("Backends: api1:8080, api2:8080, api3:8080")
	http.ListenAndServe(":8080", nil)
}