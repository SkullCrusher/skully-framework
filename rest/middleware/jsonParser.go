package middleware

import (
	"encoding/json"
	"net/http"
)

/*
	JsonParserMiddleware
	Parses the request json object to a interface.

	@param {...}   arg          - The function that gets called down the chain.

	@returns {http.ResponseWriter} w
	@returns {*http.Request}       r
*/
func JsonParserMiddleware(arg func(w http.ResponseWriter, r *http.Request, parsed interface{})) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {

		// Create a decoder we need.
		decoder := json.NewDecoder(r.Body)

		var t interface{}
		err := decoder.Decode(&t)

		// If there was a decoding error, reject.
		if err != nil{
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Make sure we close our body.
		defer r.Body.Close()

		// Process the request.
		arg(w, r, t)
	}
}
