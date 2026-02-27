package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadyz godoc
//
//	@Summary		Get readiness
//	@ID				get-readyz
//	@Tags			healthcheck
//	@Success		200	{object}		object
//	@Failure		500	{object}	object
//	@Router		  /readyz [get]
func (c *Controller) GetReadyz(ctx *gin.Context) {
	writer := ctx.Writer
	writer.WriteHeader(http.StatusOK)
}

// GetLivez godoc
//
//	@Summary		Get liveness
//	@ID				get-livez
//	@Tags			healthcheck
//	@Success		200	{object}		object
//	@Failure		500	{object}	object
//	@Router		  /livez [get]
func (c *Controller) GetLivez(ctx *gin.Context) {
	writer := ctx.Writer
	writer.WriteHeader(http.StatusOK)
}
