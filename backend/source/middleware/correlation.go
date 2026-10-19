package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/F4nk1/agentrix/backend/source/tracer"
)

func Correlation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		requestId := uuid.New().String()
		ctx := context.WithValue(r.Context(), tracer.RequestIdKey, requestId)
		w.Header().Set("X-Request-Id",requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}