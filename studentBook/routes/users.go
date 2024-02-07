package routes

import (
	"gilab.com/pragmacreviews/golang-gin-poc/handlers/users"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	router.GET("/users", getAllUsers)
	//router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	//router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
}

func getAllUsers(c *gin.Context) {
	allUsers := users.GetAllUsers(c)
	c.JSON(200, gin.H{
		"All students": allUsers,
	})
}

func createUser(c *gin.Context) {
	postUser := users.CreateUser(c)
	c.JSON(201, gin.H{
		"You posted a new student": postUser,
	})
}

func deleteUser(c *gin.Context) {
	delUser := users.DeleteUser(c)
	c.JSON(200, gin.H{
		"You deleted a user": delUser,
	})
}
