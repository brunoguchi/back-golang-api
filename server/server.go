package server

import (
	"log"

	"github.com/brunoguchi/back-golang-api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "6000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Servidor rodando na porta: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
