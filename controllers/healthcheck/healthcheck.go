package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logan1o1/go_chess_server/interfaces"
)

type HealthCheckController struct {
	service interfaces.IHealthCheckService
}

func (hcc *HealthCheckController) CheckServerHealthCheck(ctx *gin.Context) {
	check, err := hcc.service.CheckHealthCheck()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, check)
}

func NewHealthCkeckController(hcsi interfaces.IHealthCheckService) *HealthCheckController {
	return &HealthCheckController{service: hcsi}
}
