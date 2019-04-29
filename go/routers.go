/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	result := GetSecrets()

	resultJson, err := json.Marshal(result)
	handleErr(err)
	// fmt.Println(string(b))
	fmt.Fprintf(w, string(resultJson))
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"AddSecret",
		strings.ToUpper("Post"),
		"/v1/secret",
		AddSecret,
	},

	Route{
		"GetSecretByHash",
		strings.ToUpper("Get"),
		"/v1/secret/{hash}",
		GetSecretByHash,
	},
}
