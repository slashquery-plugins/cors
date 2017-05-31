package cors

import (
	"log"
	"net/http"
	"time"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().UTC().String()
		log.Println("CORS init")
		w.Header().Set("sq-cors-version", time)
		r.Header.Set("sq-cors-version", time)
		next.ServeHTTP(w, r)
	})
}
