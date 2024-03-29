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
			"subjectName": subjectName,
		}
		subjects = append(subjects, subject)
	}
	return subjects
}

func CreateSubject(c *gin.Context) []map[string]string {
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
		return nil
	}

	_, err = db.Exec("insert into subjects(name) values(?)", subjectName)
	if err != nil {
		log.Fatal("Failed to insert a subject:", err)
	}
	subjectName1 := make(map[string]string)
	subjectName1["name"] = subjectName
	return []map[string]string{subjectName1}
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
