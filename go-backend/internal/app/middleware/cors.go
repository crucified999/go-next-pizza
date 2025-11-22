package middleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		
		if origin == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		} else if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		
		w.Header().Set("Access-Control-Max-Age", "86400")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func CORSOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
					origin := r.Header.Get("Origin")
					
					if origin == "http://localhost:3000" {
							w.Header().Set("Access-Control-Allow-Origin", origin)
							w.Header().Set("Access-Control-Allow-Credentials", "true")
					} else if origin != "" {
							w.Header().Set("Access-Control-Allow-Origin", origin)
					} else {
							w.Header().Set("Access-Control-Allow-Origin", "*")
					}
					
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
					w.Header().Set("Access-Control-Max-Age", "86400")
					
					w.WriteHeader(http.StatusNoContent)
					return
			}
			
			next.ServeHTTP(w, r)
	})
}