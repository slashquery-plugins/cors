package cors

import (
	"log"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CORS init")
		w.Header().Set("sq-cors-version", "1.0.0")
		r.Header.Set("sq-cors-version", "1.0.0")
		next.ServeHTTP(w, r)
	})
}
