package helpers

import (
	"encoding/json"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request)(bool, interface{}){

	// Create a decoder we need.
	decoder := json.NewDecoder(r.Body)

	var t interface{}
	err := decoder.Decode(&t)

	// If there was a decoding error, reject.
	if err != nil{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false, t
	}

	// Make sure we close our body.
	defer r.Body.Close()


	return true, t
}