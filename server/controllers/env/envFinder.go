package env

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "template.clevecord.me/server/apimodels/env"
	"template.clevecord.me/server/constants"
	service "template.clevecord.me/server/services/env"
	"template.clevecord.me/server/utilities"
)

// GetEnv returns the value of the corresponding environment variable if it exists.
func GetEnv(c *gin.Context) {
	envKey := c.Param("env")
	if len(envKey) <= 0 {
		errorResponse, _ := utilities.CraftErrorResponse(utilities.ErrorResponse{
			Message:     "env key required",
			Endpoint:    c.Request.URL.String(),
			RequestType: c.Request.Method,
		})
		c.Data(http.StatusInternalServerError, constants.ContentTypeJSON, errorResponse)
		return
	}

	value, err := service.LookupEnv(envKey)
	if err != nil {
		errorResponse, _ := utilities.CraftErrorResponse(utilities.ErrorResponse{
			Message:     err.Error(),
			Endpoint:    c.Request.URL.String(),
			RequestType: c.Request.Method,
		})
		c.Data(http.StatusInternalServerError, constants.ContentTypeJSON, errorResponse)
		return
	}

	response := model.EnvGetResponse{
		Key:   envKey,
		Value: value,
	}

	jsonResponse, err := utilities.CraftJSONResponse(response, c.Request.URL.String(), c.Request.Method)
	var returnCode = http.StatusOK
	if err != nil {
		returnCode = http.StatusInternalServerError
	}
	c.Data(returnCode, constants.ContentTypeJSON, jsonResponse)
}
