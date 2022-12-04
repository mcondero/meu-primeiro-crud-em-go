package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mauriciocondero/meu-primeiro-crud-em-go/src/configuration/rest_err"
	"k8s.io/apiserver/pkg/admission/plugin/webhook/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err)
		)
	}

}
