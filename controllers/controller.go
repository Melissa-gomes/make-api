package controllers

import (
	"fmt"
	"project-api/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service services.Service
}

func NewController(service services.Service) Controller {
	return Controller{Service: service}
}

func (ctl Controller) Get(ctx *gin.Context) {
	fmt.Println("óh eu aqui")
	res := ctl.Service.ServiceGet()
	fmt.Println(res)
}
