package middleware

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
)

//DocumentationHandler Expose Redoc documentation on /docs endpoint
func DocumentationHandler(serverURL string, path string) http.Handler {
	handler := http.NotFoundHandler()
	if os.Getenv("EXPOSE_DOC") == "true" {
		handler = middleware.Redoc(middleware.RedocOpts{
			Path:    path,
			SpecURL: serverURL + "/swagger",
		}, handler)
	}
	return handler
}

//SwaggerHandler Expose local swagger file on /swagger endpoint
func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("EXPOSE_DOC") == "true" {
		r.Header.Add("Content-Type", "application/json")
		data, err := ioutil.ReadFile("swagger.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write(data)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
