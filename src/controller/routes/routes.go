package routes

import (
	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:useremail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId")
}
