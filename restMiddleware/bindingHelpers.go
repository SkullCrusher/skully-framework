package restMiddleware

import (
	"fmt"
	"net/http"
)

/*
	forceHeaders
	Handle forcing the headers for the rest calls to bypass cors.

	@param {http.ResponseWriter} w - The request we want to force the headers on.

	@return null
*/
func forceHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, namespace")
	w.Header().Set("Content-Type", "application/json")
}

/*
	preProcess
	preprocess every network request with common code.

	@params {http.ResponseWriter} w
	@params {*http.Request} r

	@returns {bool}
	@returns {string}
*/
func preProcess(w http.ResponseWriter, r *http.Request)(bool, string){

	forceHeaders(w)

	// Block option requests from being processed.
	if r.Method == "OPTIONS" {
		_, _ = fmt.Fprintf(w, "")
		return false, ""
	}

	// Get the namespace if provided.
	namespaceHeader := r.Header.Get("namespace")

	return true, namespaceHeader
}
