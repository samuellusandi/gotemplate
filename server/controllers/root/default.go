package root

import (
	"net/http"

	rootmodels "template.clevecord.me/server/apimodels/root"
	"template.clevecord.me/server/constants"
	"template.clevecord.me/server/utilities"

	"github.com/gin-gonic/gin"
)

// GetDefault returns the default handler for requests
func GetDefault(c *gin.Context) {
	var appVersion string = constants.Lookup(constants.AppVersionEnvKey)
	var appEnvironment string = constants.Lookup(constants.AppEnvironmentEnvKey)
	var appDescription string = constants.Lookup(constants.AppDescriptionEnvKey, "No description provided.")

	var appResponse interface{} = rootmodels.App{
		Version:     appVersion,
		Description: appDescription,
		Environment: appEnvironment,
	}

	jsonResponse, err := utilities.CraftJSONResponse(appResponse, c.Request.URL.String(), c.Request.Method)
	var returnCode = http.StatusOK
	if err != nil {
		returnCode = http.StatusInternalServerError
	}
	c.Data(returnCode, constants.ContentTypeJSON, jsonResponse)
}
