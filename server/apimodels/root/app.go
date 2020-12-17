package apimodels

// App represents the application JSON Model for REST API use.
type App struct {
	Description string `json:"app_description"`
	Version     string `json:"app_version"`
	Environment string `json:"app_environment"`
}
