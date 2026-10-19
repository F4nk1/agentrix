package middleware

import (
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
    
    "github.com/F4nk1/agentrix/backend/source/cache"
    "github.com/F4nk1/agentrix/backend/source/common"
    "github.com/F4nk1/agentrix/backend/source/model"
    "github.com/F4nk1/agentrix/backend/source/tracer"
)
    
func Cache(c cache.Cache) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            if r.Method != http.MethodGet {
                next.ServeHTTP(w, r)
                return
            }
    
            key := r.URL.Path
            if claims, ok := ctx.Value(common.UserContextKey).(*model.Claims); ok {
                key = fmt.Sprintf("%d:%s", claims.UserId, key)
            }
    
            if cached, found := c.Get(key); found {
                tracer.Debugf(ctx, "Cached response served for %s", key)
                w.Header().Set("Content-Type", common.ContentTypeJSON)
                w.Write(cached.([]byte))
                return
            }
    
            recorder := cache.NewRecorder(w)
            next.ServeHTTP(recorder, r)
            c.Set(key, recorder.Body, cache.DefaultTTL)
        })
    }
}