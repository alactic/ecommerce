package authentication

import (
	"log"
	"net/http"
	"os"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout) // logs go to Stderr by default
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler
	})
}
