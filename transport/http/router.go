package http

func (s *Server) InitRoutes() {
	v1 := s.router.Group("transactions")
	v1.POST("", s.handler.ServeTransaction)
	v1.GET("/history", s.handler.GetTransaction)
	v1.GET("/history/statistics", s.handler.GetStatistics)
}
