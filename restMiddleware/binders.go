package restMiddleware

import (
	"github.com/gorilla/mux"
	"net/http"
)


/*
	bindRoute
	handles binding with middleware that applies to everything

	@param {*mux.Router} r       - Router pointer.
	@param {string}      route   - The route we want to bind.
	@param {...}         arg     - The callback function to handle the request.
	@param {string       methods - The http method that can be used to access the route.

	@returns null
*/
func bindRoute(r *mux.Router, route string, handler func(w http.ResponseWriter, r *http.Request), methods string){
	r.HandleFunc(route, handler).Methods(methods)
	r.HandleFunc(route, handler).Methods("OPTIONS")
}

/*
	handleNoAuth
	Handles a route that is public and unrestricted.

	@param {...} arg - Callback for the request.

	@returns {http.ResponseWriter}
	@returns {r *http.Request}
*/
func handleNoAuth(arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Preprocess the request.
		preProcess, namespaceHeader := preProcess(w, r)

		// If we shouldn't continue.
		if preProcess == false{
			return
		}

		arg(w, r, "", namespaceHeader)
	}
}

/*
	handleAuth
	Handles a route that is authentication restricted.

	@param {...} arg - Callback for the request.

	@returns {http.ResponseWriter}
	@returns {r *http.Request}
*/
func handleAuth(arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string), authFunc func(w http.ResponseWriter, r *http.Request)(bool, string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Preprocess the request.
		preProcess, namespaceHeader := preProcess(w, r)

		// If we shouldn't continue.
		if preProcess == false {
			return
		}

		// Validate the authentication.
		validUser, userId := authFunc(w, r)

		// If the user shouldn't be able to access this resource.
		if validUser == false{
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		arg(w, r, userId, namespaceHeader)
	}
}
