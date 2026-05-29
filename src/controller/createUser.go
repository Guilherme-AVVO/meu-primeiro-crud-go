package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

