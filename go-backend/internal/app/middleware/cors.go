package middleware

import (
	"net/http"
)

// CORS middleware для разрешения доступа всем клиентам
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем все источники
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Разрешаем все методы
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		
		// Разрешаем все заголовки
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		
		// Разрешаем кеширование preflight запросов на 24 часа
		w.Header().Set("Access-Control-Max-Age", "86400")
		
		// Обрабатываем preflight запросы
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}