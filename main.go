package main

import (
	docs "github.com/cold-vodka/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")

	v1.GET("/", HelloWorld)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
