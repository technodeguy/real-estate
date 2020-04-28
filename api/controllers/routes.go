package controllers

import (
	"github.com/technodeguy/real-estate/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home route
	s.router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.HealthCheck)).Methods("GET")

	// User routes
	s.router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.router.HandleFunc("/users/presigned_url", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.GetPresignedUrl))).Methods("POST")
	s.router.HandleFunc("/users/avatar", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.SaveUserAvatar))).Methods("PUT")
	s.router.HandleFunc("/users/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Estate routes
	s.router.HandleFunc("/estates/cities/{city_id}", middlewares.SetMiddlewareJSON(s.GetEstatesByCityId)).Methods("GET")
	s.router.HandleFunc("/estates", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.CreateEstate))).Methods("POST")
	s.router.HandleFunc("/estates/sell", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.SellEstate))).Methods("PUT")

}
