package subjects

import (
	"gilab.com/pragmacreviews/golang-gin-poc/core/subjects"
	"github.com/gin-gonic/gin"
)

func GetAllSubjects(c *gin.Context) []map[string]string {
	allSubjects := subjects.GetAllSubjects(c)
	return allSubjects
}

func GetSubjectByName(c *gin.Context) string {
	subject := subjects.GetSubjectByName(c)
	return subject
}

func CreateSubject(c *gin.Context) string {
	createSubject := subjects.CreateSubject(c)
	return createSubject
}

func DeleteSubject(c *gin.Context) string {
	deleteSubject := subjects.DeleteSubject(c)
	return deleteSubject
}
