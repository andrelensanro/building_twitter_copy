package middlew

import (
	"net/http"

	"github.com/andrelensanro/building_twitter_copy/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Disconnect Data Base", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
