package middleware

import (
	"fmt"
	"net/http"
)

/*
	forceHeaders
	Handle forcing the headers for the rest calls.

	@param {http.ResponseWriter} w - The request we want to force the headers on.

	@return null
*/
func ForceHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
}


/*
	NoAuthRouteWrapper
	Pre-handle a route that

	@param {func(w http.ResponseWriter, r *http.Request)} arg - The route to handle the request.

	@returns func(w http.ResponseWriter, r *http.Request)
*/
func NoAuthRouteWrapper(arg func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {

		// Force the headers.
		ForceHeaders(w)

		// Block option requests from being processed.
		if r.Method == "OPTIONS" {
			_, _ = fmt.Fprintf(w, SuccessMessageFormat, "")
			return
		}

		arg(w, r)
	}
}


/*
	AuthRouteWrapper
	Force the authentication to happen first for this route.

	@param {func} arg - The route to handle the request.
	@param {func} authFunc - The function that validates the authentication request

	@returns func(w http.ResponseWriter, r *http.Request)
*/
func AuthRouteWrapper(arg func(w http.ResponseWriter, r *http.Request, userId string), authFunc func(token string)(bool, string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Force the headers.
		ForceHeaders(w)

		// Block option requests from being processed.
		if r.Method == "OPTIONS" {
			_, _ = fmt.Fprintf(w, SuccessMessageFormat, "")
			return
		}

		// Get the authorization header.
		authenticationHeader := r.Header.Get("Authorization")

		// Validate it.
		validUser, userId := authFunc(authenticationHeader)

		if validUser == false{
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		arg(w, r, userId)
	}
}
