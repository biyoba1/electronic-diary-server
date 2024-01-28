package subjects

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB
var err error

// todo сделать универсальным метод подклюения к таблицы
func OpenDataBase() {
	db, err = sql.Open("sqlite3", "subjects.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
}

func GetAllSubjects(c *gin.Context) []map[string]string {
	OpenDataBase()
	defer db.Close()
	rows, err := db.Query("SELECT name FROM subjects")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return nil
	}
	defer rows.Close()

	var subjects []map[string]string
	for rows.Next() {
		var subjectName string
		err := rows.Scan(&subjectName)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return nil
		}
		subject := map[string]string{
			"subject": subjectName,
		}
		subjects = append(subjects, subject)
	}
	return subjects
}

func GetSubjectByName(c *gin.Context) string {
	OpenDataBase()
	defer db.Close()
	subjectType := c.Query("type")
	if subjectType == "" {
		c.JSON(400, gin.H{"error": "Missing 'type' parameter"})
		return "null"
	}

	rows, err := db.Query("SELECT name FROM subjects WHERE name = ?", subjectType)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return "null"
	}
	defer rows.Close()

	var subject string
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return ""
		}
		subject = result
	}
	return subject
}

func CreateSubject(c *gin.Context) string {
	OpenDataBase()
	defer db.Close()
	subjectName := c.Query("name")
	if len(subjectName) == 0 {
		c.JSON(400, gin.H{
			"error": "name are required",
		})
	}

	if subjectExists(subjectName) {
		c.JSON(500, gin.H{
			"error": "subject already in database",
		})
		return "null"
	}

	_, err = db.Exec("insert into subjects(name) values(?)", subjectName)
	if err != nil {
		log.Fatal("Failed to insert a subject:", err)
	}
	return subjectName
}

func subjectExists(name string) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM subjects WHERE name = (?)", name)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Failed to check if student exists:", err)
	}
	return count > 0
}

func DeleteSubject(c *gin.Context) string {
	OpenDataBase()
	defer db.Close()
	subjectName := c.Query("name")
	if len(subjectName) == 0 {
		c.JSON(400, gin.H{
			"error": "name are required",
		})
	}
	_, err := db.Exec("delete from subjects where name = (?)", subjectName)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to delete a subject",
		})
	}
	return subjectName
}
