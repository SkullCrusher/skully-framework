package helpers

import "net/http"

/*
	ForceHeaders
	Handle forcing the headers for the rest calls to bypass cors.

	@param {http.ResponseWriter} w - The request we want to force the headers on.

	@return none
*/
func ForceHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, namespace")
	w.Header().Set("Content-Type", "application/json")
}