package grades

import (
	"database/sql"
	_ "gilab.com/pragmacreviews/golang-gin-poc/core/users"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB
var db3 *sql.DB
var db2 *sql.DB
var err error

func OpenDataBase() {
	db, err = sql.Open("sqlite3", "combined.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	db2, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	db3, err = sql.Open("sqlite3", "subjects.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
}

func GetAllGrades(c *gin.Context) ([]map[string]string, []map[string]string, []map[string]string, []map[string]string) {
	db, err = sql.Open("sqlite3", "combined.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	rows, err := db.Query("select name, second_name, patronymic, subject_name, grade, date1 from combined")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return nil, nil, nil, nil
	}
	defer rows.Close()

	var users []map[string]string
	var subjects []map[string]string
	var grades []map[string]string
	var dates []map[string]string

	for rows.Next() {
		var secondName, name, patronymic, subject_name, grade1, date1 string
		err := rows.Scan(&secondName, &name, &patronymic, &subject_name, &grade1, &date1)
		if err != nil {
			return nil, nil, nil, nil
		}

		user := map[string]string{
			"Second name": secondName,
			"Name":        name,
			"Patronymic":  patronymic,
		}

		subject := map[string]string{
			"Subject name": subject_name,
		}

		grade := map[string]string{
			"Grade": grade1,
		}

		date := map[string]string{
			"Date": date1,
		}

		users = append(users, user)
		subjects = append(subjects, subject)
		grades = append(grades, grade)
		dates = append(dates, date)
	}
	return users, subjects, grades, dates
}

func PostGrades(c *gin.Context) ([]map[string]string, []map[string]string, []map[string]string) {
	OpenDataBase()
	defer db.Close()
	defer db2.Close()
	defer db3.Close()

	//student
	userName := c.Query("name")
	userSecondName := c.Query("secondName")
	userPatronymic := c.Query("patronymic")

	//subject
	subjectName := c.Query("subjectName")

	//grade
	grade := c.Query("grade")
	date := c.Query("date")

	//вносим наши данные в саму таблицу
	_, err := db.Exec("insert into combined(name, second_name, patronymic, subject_name, grade, date1) values (?,?,?,?,?,?)", userName, userSecondName, userPatronymic, subjectName, grade, date)
	if err != nil {
		log.Fatal("Failed to insert information to database:", err)

	}
	user := make(map[string]string)
	user["second name"] = userSecondName
	user["name"] = userName
	user["patronymic"] = userPatronymic
	subject := make(map[string]string)
	subject["subject"] = subjectName
	grade1 := make(map[string]string)
	grade1["grade"] = grade
	return []map[string]string{user}, []map[string]string{subject}, []map[string]string{grade1}
}
