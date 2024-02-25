package users

import (
	"gilab.com/pragmacreviews/golang-gin-poc/core/users"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) []map[string]string {
	allUsers := users.GetAllStudents(c)
	return allUsers
}

func CreateUser(c *gin.Context) []map[string]string {
	postUser := users.CreateUser(c)
	return postUser
}

func DeleteUser(c *gin.Context) []map[string]string {
	deleteUser := users.DeleteUser(c)
	return deleteUser
}
