package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajajamal/Runners-Board/models"
)

type RunnersController struct {
	runnersService *services.RunnersService
}

func NewRunnersController(runnersService *services.RunnersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
	}
}
func (rh RunnersController) CreateRunner(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Error while unmarshalling"+"create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := rh.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
func (rh RunnersController) UpdateRunnerRunner(ctx *gin.Context) {

}
func (rh RunnersController) DeleteRunnerRunner(ctx *gin.Context) {

}
func (rh RunnersController) GetRunnerRunner(ctx *gin.Context) {

}
func (rh RunnersController) GetRunnersBatchRunner(ctx *gin.Context) {

}
