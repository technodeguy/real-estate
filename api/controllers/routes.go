package controllers

func (s *Server) initializeRoutes() {
	s.router.HandleFunc("/", s.Home).Methods("GET")
}
