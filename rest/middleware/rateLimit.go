package middleware

import (
	"../helpers"
	"../logic"
	"golang.org/x/time/rate"
	"net/http"
)

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
	var limiter = logic.NewIPRateLimiter(rate.Limit(refillRate), defaultStock)

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the ip address from cloudflare or raw ip.
		ipAddress := helpers.GetIP(r)

		limiter := limiter.GetLimiter(ipAddress) // r.RemoteAddr)

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		arg(w, r)
	}
}
