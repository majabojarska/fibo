package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetFibonacci(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HELLO THERE FROM THE CONTROLLER")
}
