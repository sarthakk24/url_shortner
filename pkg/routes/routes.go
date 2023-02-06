package routes

import (
	"urlShortner/pkg/api/controllers/authentication"
	test "urlShortner/pkg/api/controllers/test"

	"github.com/gin-gonic/gin"
)

func AllRoutes() {
	router := gin.Default()
	router.GET("/test", test.TestMongo)

	router.POST("/signup", authentication.SignUp)
	router.POST("/signin", authentication.SignIn)

	router.Run("localhost:8080")
}
