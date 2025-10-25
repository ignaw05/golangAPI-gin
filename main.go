package main

import (
	"Gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetUpRouter(r)

	r.Run(":8080")
}