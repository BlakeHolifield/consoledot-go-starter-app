package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(apiPath string) *gin.Engine {
	router := gin.Default()
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath)
	return router
}
