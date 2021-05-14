package restMiddleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
	BindNoAuthRoutes
	Binds a new public route that doesn't require auth.

	@param {*mux.Router} r            - Router pointer.
	@param {string}      route        - The route we want to bind.
	@param {...}         arg          - The callback function to handle the request.
	@param {int64}       refillRate   - The how many requests are added back to the rate limiter per second.
	@param {int}         defaultStock - The starting value for the rate limiter pool.
	@param {string       methods      - The http method that can be used to access the route.

	@returns null
*/
func BindNoAuthRoutes(r *mux.Router, route string, arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string), refillRate int64, defaultStock int, methods string){

	// If we don't have a refill rate then don't attach a rate limiter.
	if refillRate == 0{
		bindRoute(r, route, handleNoAuth(arg), methods)
	}else{
		bindRoute(r, route, LimitMiddleware(handleNoAuth(arg), refillRate, defaultStock), methods)
	}
}

/*
	BindAuthRoute
	Binds a new public route that requires auth.

	@param {*mux.Router} r            - Router pointer.
	@param {string}      route        - The route we want to bind.
	@param {...}         arg          - The callback function to handle the request.
	@param {int64}       refillRate   - The how many requests are added back to the rate limiter per second.
	@param {int}         defaultStock - The starting value for the rate limiter pool.
	@param {...}         auth         - The function to check if the request is valid or not.
	@param {string       methods      - The http method that can be used to access the route.

	@returns null
*/
func BindAuthRoute(r *mux.Router, route string, arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string), refillRate int64, defaultStock int, auth func(w http.ResponseWriter, r *http.Request)(bool, string), methods string){

	// If we don't have a refill rate then don't attach a rate limiter.
	if refillRate == 0{
		bindRoute(r, route, handleAuth(arg, auth), methods)
	}else{
		bindRoute(r, route, LimitMiddleware(handleAuth(arg, auth), refillRate, defaultStock), methods)
	}
}