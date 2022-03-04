package middlewares

import "net/http"
import "log"

func CheckBody(next http.HandlerFunc) http.HandlerFunc {
	return func (res http.ResponseWriter, req *http.Request) {
		log.Printf("Middleware CheckBody di jalankan")
		next.ServeHTTP(res, req)
	}
}