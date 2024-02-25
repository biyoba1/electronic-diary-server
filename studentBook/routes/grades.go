package routes

import (
	"gilab.com/pragmacreviews/golang-gin-poc/handlers/grades"
	"github.com/gin-gonic/gin"
)

func GradesRoutes(router *gin.Engine) {
	router.GET("/grades", getAllGrades)
	router.POST("/grades", postGrades)
	router.DELETE("/grades/:id", deleteUser)
}

func getAllGrades(c *gin.Context) {
	student, subject, grades, date := grades.GetAllGrades(c)
	c.JSON(200, gin.H{
		"all_students": student,
		"subject":      subject,
		"grades":       grades,
		"date":         date,
	})
}

func postGrades(c *gin.Context) {
	student, subject, grades := grades.PostGrades(c)
	c.JSON(201, gin.H{
		"student": student,
		"subject": subject,
		"grades":  grades,
	})
}

//func createGrades(c *gin.Context) {
//	student1, subject1, grades26 := grades.CreateGrades(c)
//	c.JSON(200, gin.H{
//		"student": student1,
//		"subject": subject1,
//		"grades":  grades26,
//	})
//}

//func createUser(c *gin.Context) {
//	student1, subject1, grade1 := grades.CreateUser(c)
//	c.JSON(200, gin.H{
//		"student": student1,
//		"subject": subject1,
//		"grade":   grade1,
//	})
//}
