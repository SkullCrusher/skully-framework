package restMiddleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

// https://pliutau.com/rate-limit-http-requests/

// IPRateLimiter .
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter .
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map, using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}

/*
	limitMiddleware
	Attaches a rate limiter to network requests (5 requests per second cap).

	@param {...}   arg          - The function that gets called down the chain.
	@param {int64} refillRate   - The how many requests are added back to the rate limiter per second.
	@param {int}   defaultStock - The starting value for the rate limiter pool.

	@returns {http.ResponseWriter} w
	@returns {*http.Request}       r
*/
func LimitMiddleware(arg func(w http.ResponseWriter, r *http.Request), refillRate int64, defaultStock int) func(w http.ResponseWriter, r *http.Request){

	// Setup the rate limiter for network calls.
	var limiter = NewIPRateLimiter(rate.Limit(refillRate), defaultStock)

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the ip address from cloudflare or raw ip.
		ipAddress := GetIP(r)

		limiter := limiter.GetLimiter(ipAddress) // r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		arg(w, r)
	}
}
