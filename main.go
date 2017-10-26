package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/phanvudinh/go-keycloak-adapter"
	"github.com/phanvudinh/go-practice-app/entity"
)

func generationEntities(connection *gorm.DB) {
	if !connection.HasTable(&entity.User{}) {
		connection.AutoMigrate(&entity.User{})
	} else {
		fmt.Println("entity exsisted")
	}
}

func main() {
	connection, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=sample sslmode=disable password=postgres")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Start Generate Entity")
		generationEntities(connection)
		fmt.Printf("End Generate Entity")
	}

	defer connection.Close()

	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var users []entity.User

		if err = c.BindJSON(&users); err == nil {
			i := 0
			length := len(users)
			for i < length {
				fmt.Println("create " + users[i].Name)
				connection.Create(&entity.User{Name: users[i].Name, Age: users[i].Age, Address: users[i].Address, Avatar: users[i].Avatar})
				i++
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.POST("/user", func(c *gin.Context) {
		var user entity.User
		if err = c.BindJSON(&user); err == nil {
			connection.Create(&user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.GET("/user", func(c *gin.Context) {
		var users []entity.User
		connection.Find(&users)
		c.JSON(200, users)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		var user entity.User
		userID, _ := strconv.Atoi(c.Param("id"))
		connection.Find(&user, userID)
		c.JSON(200, user)
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		var user entity.User
		userID, _ := strconv.Atoi(c.Param("id"))
		connection.Find(&user, userID)
		connection.Unscoped().Delete(&user)
	})

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
