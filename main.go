package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phanvudinh/go-keycloak-adapter"
)

func main() {
	r := gin.Default()
	r.GET("/userinfo", func(c *gin.Context) {
		// get access_token from http request
		goKeycloakAdapter.GetUserInfo(c.Request)
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/authorization", func(c *gin.Context) {
		// get access_token from http request
		goKeycloakAdapter.IsAuthorized(c.Request)
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
