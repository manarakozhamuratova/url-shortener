package transport

import echoSwagger "github.com/swaggo/echo-swagger"

func (s *Server) SetupRoutes() {
	s.App.GET("/swagger/*", echoSwagger.WrapHandler)
	s.App.POST("/shortener", s.handler.CreateShortUrl)
	s.App.GET("/shortener", s.handler.GetURLs)
	s.App.GET("/:link", s.handler.GetLongURL)
	s.App.DELETE("/:link", s.handler.DeleteURL)
	s.App.GET("/stats/:link", s.handler.GetStats)
	s.App.GET("/health", s.handler.HealthCheck)
}
