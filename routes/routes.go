package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID int "json:'id'"
	Nombre string "json:'nombre'"
	Email string "json:'email'"
}

var usuarios []Usuario
func SetUpRouter(r *gin.Engine){

	r.GET("/", func (c *gin.Context) {
		// c.String(200, "Hola mundo")
		c.JSON(http.StatusOK, usuarios)
	})

	r.GET("/saludo/:nombre", func(c *gin.Context){
		nombre:=c.Param("nombre")
		c.String(http.StatusOK,"!Hola %s!", nombre)
	})

	r.POST("/usuarios", func (c *gin.Context){
		var nuevoUsuario Usuario

		if err:= c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Error al decodificar el json"})
			return
		}

		if (nuevoUsuario.Nombre == "") || (nuevoUsuario.Email == "") || (nuevoUsuario.ID == 0) {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Todos los datos son obligatorios"})
		}

		usuarios = append(usuarios,nuevoUsuario)
		c.JSON(http.StatusOK, gin.H{"mensaje":"Usuario registrado correactamente"})
	})
}