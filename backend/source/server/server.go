package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/F4nk1/agentrix/backend/source/auth"
	"github.com/F4nk1/agentrix/backend/source/cache"
	"github.com/F4nk1/agentrix/backend/source/middleware"
	"github.com/F4nk1/agentrix/backend/source/service"
)

type Server struct {
	Handler			http.Handler
	Service			service.Service
	SessionStore	*sessions.CookieStore
	Cache			cache.Cache
	Auth			*auth.Auth
}

func NewServer(service service.Service) *Server {
	s := &Server{
		Service:		service,
		SessionStore:	sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET"))),
		Cache:			cache.NewSimpleCache(),
		Auth:			auth.NewAuth(os.Getenv("ACCESS_SECRET"),os.Getenv("REFRESH_SECRET")),
	}
	s.Handler = s.buildHandler()
	return s
}

func (s *Server) buildHandler() http.Handler {
	router := mux.NewRouter()
	
	router.Use(middleware.Correlation)
	router.Use(middleware.Cors)

	api := router.PathPrefix("/api/v1").Subrouter()

	protected := api.NewRoute().Subrouter()
	protected.Use(middleware.Auth(s.Auth))
	protected.Use(middleware.Cache(s.Cache))
	return router
}
