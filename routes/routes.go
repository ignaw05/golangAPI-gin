package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine){

	r.GET("/", func (c *gin.Context) {
		// c.String(200, "Hola mundo")
		c.JSON(http.StatusOK, gin.H{
			"message":"Hola mundillo",
	})})

	r.GET("/saludo/:nombre", func(c *gin.Context){
		nombre:=c.Param("nombre")
		c.String(http.StatusOK,"!Hola %s!", nombre)
	})
}