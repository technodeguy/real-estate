package controllers

import (
	"github.com/technodeguy/real-estate/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home route
	s.router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.HealthCheck)).Methods("GET")

	//User routes
	s.router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.router.HandleFunc("/users/presigned_url", middlewares.SetMiddlewareJSON(s.GetPresignedUrl)).Methods("POST")
	s.router.HandleFunc("/users/avatar", middlewares.SetMiddlewareJSON(s.SaveUserAvatar)).Methods("PUT")
	s.router.HandleFunc("/users/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

}
