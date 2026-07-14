package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// CanaryProxy handles intelligent traffic splitting.
type CanaryProxy struct {
	V1Proxy *httputil.ReverseProxy
	V2Proxy *httputil.ReverseProxy
	WeightV2 int // Percentage of traffic to send to V2 (0-100)
}

func NewCanaryProxy(v1Addr, v2Addr string, weightV2 int) *CanaryProxy {
	v1URL, _ := url.Parse(v1Addr)
	v2URL, _ := url.Parse(v2Addr)

	return &CanaryProxy{
		V1Proxy:  httputil.NewSingleHostReverseProxy(v1URL),
		V2Proxy:  httputil.NewSingleHostReverseProxy(v2URL),
		WeightV2: weightV2,
	}
}

func (p *CanaryProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Canary Logic: Decision based on traffic weight
	if rand.Intn(100) < p.WeightV2 {
		log.Printf("Routing request to V2: %s", r.URL.Path)
		p.V2Proxy.ServeHTTP(w, r)
	} else {
		log.Printf("Routing request to V1: %s", r.URL.Path)
		p.V1Proxy.ServeHTTP(w, r)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Example: 20% of traffic goes to V2 (Canary)
	proxy := NewCanaryProxy("http://localhost:8081", "http://localhost:8082", 20)
	
	log.Println("Canary Proxy running on :8080...")
	http.ListenAndServe(":8080", proxy)
}
