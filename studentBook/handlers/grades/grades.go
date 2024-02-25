package grades

import (
	"gilab.com/pragmacreviews/golang-gin-poc/core/grades"
	"github.com/gin-gonic/gin"
)

func GetAllGrades(c *gin.Context) ([]map[string]string, []map[string]string, []map[string]string, []map[string]string) {
	student, subject, grades, date := grades.GetAllGrades(c)
	return student, subject, grades, date
}

func PostGrades(c *gin.Context) ([]map[string]string, []map[string]string, []map[string]string) {
	student, subject, grades := grades.PostGrades(c)
	return student, subject, grades
}

//func CreateGrades(c *gin.Context) ([]map[string]string, []map[string]string, []map[string]string) {
//	student1, subject1, grades26 := subjects.CreateGrades(c)
//	return student1, subject1, grades26
//}
