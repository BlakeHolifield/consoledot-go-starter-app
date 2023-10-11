package main

import (
	"github.com/RedHatInsights/${{ values.project_name }}/config"
	"github.com/RedHatInsights/${{ values.project_name }}/docs"
	"github.com/RedHatInsights/${{ values.project_name }}/metrics"
	"github.com/RedHatInsights/${{ values.project_name }}/providers"
	"github.com/RedHatInsights/${{ values.project_name }}/routes"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	conf            = config.Load()
	providerManager providers.Providers
	apiPath         = conf.GetApiPath()
)

// main godoc
// @title ConsoleDot Go Starter App API
// @version         1.0
// @description     API Docs for ConsoleDot Go Starter App API
// @contact.name   	Adam Drew
// @contact.email  	addrew@redhat.com
// @BasePath  		/api/${{ values.project_name }}-api/v1
func main() {
	// Initialize logging
	initLogging()

	var providerCloseFunc func(providers.Providers)
	providerManager, providerCloseFunc = providers.Init(conf)
	defer providerCloseFunc(providerManager)

	// Serve the prometheus metrics
	go metrics.Serve(conf)
	// Set up the Gin router
	router := routes.SetupRouter(apiPath, providerManager)
	// Set up the OpenAPI docs
	initAPIDocs(router)
	// Run the router
	router.Run(conf.RouterBindAddress())
}

// initLogging sets up the logging
func initLogging() {
	// Set the default log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting ConsoleDot Go Starter App API")
}

// initAPIDocs sets up the swagger (openAPI) docs
func initAPIDocs(router *gin.Engine) {
	swaggerRoute := apiPath + "/api-docs/*any"
	// Example of editing the OpenAPI info programatically
	docs.SwaggerInfo.Host = conf.RouterBindAddress()
	docs.SwaggerInfo.BasePath = swaggerRoute
	// Serve out the API Docs
	router.GET(swaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
