package middleware

import (
	"net/http"
)

type CORSMiddleware struct{}

func NewCORS() *CORSMiddleware {
	return &CORSMiddleware{}
}

// SetCORS 设置响应头
func (c *CORSMiddleware) SetCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, AccessToken, Token, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next(w, r)
	})
}
