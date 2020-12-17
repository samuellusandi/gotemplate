package routes

import (
	"net/http"

	"template.clevecord.me/server/constants"
	envController "template.clevecord.me/server/controllers/env"
	defaultControllers "template.clevecord.me/server/controllers/root"
	"template.clevecord.me/server/utilities"

	"github.com/gin-gonic/gin"
)

// InitRoutes provide the initialization of the routes exposed by this service.
func InitRoutes(e *gin.Engine) {
	// Default Routes
	e.GET(APIPrefix+DefaultPrefix, defaultControllers.GetDefault)

	// Env Routes
	e.GET(APIPrefix+EnvPrefix+"/", envController.GetEnv)
	e.GET(APIPrefix+EnvPrefix+"/:env", envController.GetEnv)

	// Not Found
	e.NoMethod(invalidMethod)
	e.NoRoute(notFound)
}

func notFound(c *gin.Context) {
	errorResponse, _ := utilities.CraftErrorResponse(utilities.ErrorResponse{
		Message:     "route not found",
		Endpoint:    c.Request.URL.String(),
		RequestType: c.Request.Method,
	})
	c.Data(http.StatusNotFound, constants.ContentTypeJSON, errorResponse)
}

func invalidMethod(c *gin.Context) {
	errorResponse, _ := utilities.CraftErrorResponse(utilities.ErrorResponse{
		Message:     "method " + c.Request.Method + " does not exist for endpoint " + c.Request.URL.String(),
		Endpoint:    c.Request.URL.String(),
		RequestType: c.Request.Method,
	})
	c.Data(http.StatusMethodNotAllowed, constants.ContentTypeJSON, errorResponse)
}
