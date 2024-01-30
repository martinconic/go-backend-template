package api

func (s *Server) InitializeRoutes() {
	s.Router.GET("/api/status", GetStatusInfo)

}
