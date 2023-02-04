package routes

import (
	test "urlShortner/pkg/api/controllers"

	"github.com/gin-gonic/gin"
)

func AllRoutes() {
	router := gin.Default()
	router.GET("/test" , test.Test)
	router.Run("localhost:8080")
}