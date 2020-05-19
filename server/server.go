package server

import "github.com/gin-gonic/gin"

func createServer() *gin.Engine {
	server := gin.Default()
	server = Routes(server)
	return server
}

func StartServer() {
	server := createServer()
	server.Run()
}
