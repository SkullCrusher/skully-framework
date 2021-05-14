package helpers

import (
	"fmt"
	"net/http"
)


/*
	PreProcess
	preprocess every network request with common code.

	@params {http.ResponseWriter} w
	@params {*http.Request} r

	@returns {bool}
	@returns {string}
*/
func PreProcess(w http.ResponseWriter, r *http.Request)(bool, string){

	ForceHeaders(w)

	// Block option requests from being processed.
	if r.Method == "OPTIONS" {
		_, _ = fmt.Fprintf(w, "")
		return false, ""
	}

	// Get the namespace if provided.
	namespaceHeader := r.Header.Get("namespace")

	return true, namespaceHeader
}
