package routes

import (
	"gilab.com/pragmacreviews/golang-gin-poc/handlers/subjects"
	"github.com/gin-gonic/gin"
)

func SubjectsRoutes(router *gin.Engine) {
	router.GET("/subject", getAllSubjects)
	router.GET("/subject/:name", getSubjectByName)
	router.POST("/subject", createSubject)
	router.DELETE("/subject/:name", deleteSubject)
}

func getAllSubjects(c *gin.Context) {
	allSubjects := subjects.GetAllSubjects(c)
	c.JSON(200, gin.H{
		"all subjects": allSubjects,
	})
}

func getSubjectByName(c *gin.Context) {
	subject := subjects.GetSubjectByName(c)
	c.JSON(200, gin.H{
		"subject": subject,
	})
}

func createSubject(c *gin.Context) {
	cSubject := subjects.CreateSubject(c)
	c.JSON(200, gin.H{
		"You posted a new subject": cSubject})
}

func deleteSubject(c *gin.Context) {
	delSubject := subjects.DeleteSubject(c)
	c.JSON(200, gin.H{
		"You delete a subject": delSubject,
	})
}
