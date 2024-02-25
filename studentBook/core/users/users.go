package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

var db *sql.DB
var err error

func userExist(userName string, userSecondName string, userPatronymic string) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ? AND second_name = ? AND patronymic = ?", userName, userSecondName, userPatronymic)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Failed to check if student exists:", err)
	}
	return count > 0
}

func OpenDataBase() {
	db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
}

func GetAllStudents(c *gin.Context) []map[string]string {
	OpenDataBase()
	defer db.Close()
	rows, err := db.Query("select second_name, name, patronymic from users")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return nil
	}
	defer rows.Close()
	var users []map[string]string
	for rows.Next() {
		var secondName, name, patronymic string
		err := rows.Scan(&secondName, &name, &patronymic)
		if err != nil {
			return nil
		}
		user := map[string]string{
			"secondName": secondName,
			"name":       name,
			"patronymic": patronymic,
		}
		users = append(users, user)
	}
	return users
}

func CreateUser(c *gin.Context) []map[string]string {
	OpenDataBase()
	defer db.Close()
	userName := c.Query("name")
	userSecondName := c.Query("secondName")
	userPatronymic := c.Query("patronymic")

	if len(userName) == 0 || len(userSecondName) == 0 || len(userPatronymic) == 0 {
		c.JSON(400, gin.H{
			"error": "name/second name/patronymic are required",
		})
	}
	if userExist(userName, userSecondName, userPatronymic) {
		c.JSON(500, gin.H{
			"error": "user already in database",
		})
		return nil
	}

	_, err := db.Exec("insert into users(name, second_name, patronymic) values (?,?,?)", userName, userSecondName, userPatronymic)
	if err != nil {
		log.Fatal("Failed to insert information to database:", err)

	}
	user := make(map[string]string)
	user["second name"] = userSecondName
	user["name"] = userName
	user["patronymic"] = userPatronymic
	return []map[string]string{user}
}

func DeleteUser(c *gin.Context) []map[string]string {
	OpenDataBase()
	defer db.Close()
	userName := c.Query("name")
	userSecondName := c.Query("secondName")
	userPatronymic := c.Query("patronymic")
	if len(userName) == 0 || len(userSecondName) == 0 || len(userPatronymic) == 0 {
		c.JSON(400, gin.H{
			"error": "name/second name/patronymic are required",
		})
	}
	_, err := db.Exec("DELETE FROM users WHERE name = ? AND second_name = ? AND patronymic = ?", userName, userSecondName, userPatronymic)
	if err != nil {
		log.Fatal("Failed to insert information to database:", err)
		return nil
	}

	user := make(map[string]string)
	user["second name"] = userSecondName
	user["name"] = userName
	user["patronymic"] = userPatronymic
	return []map[string]string{user}
}
