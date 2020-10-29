package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/plapi/config"
	"github.com/plapi/models"
	"log"
)

func CreateSubjectTable() {
	db := config.Connect()

	result, err := db.Exec("CREATE TABLE subject (id SERIAL, name TEXT, teacher TEXT, hours TEXT )")

	if err != nil {
		log.Printf("POSTGRES-ERROR: Cannot CREATE TABLE subject, Reason: %v\n", err)
	}
	if result != nil {
		log.Printf("Table Subject created")
	}
}

func InsertSubject(context *gin.Context) {
	db := config.Connect()
	params := models.Subject{}
	context.Bind(&params)

	if params.Name != "" && params.Teacher != "" && params.Hour != "" {
		subject := &models.Subject{"0", params.Name, params.Teacher, params.Hour}

		err := db.QueryRow("INSERT INTO subject (name, teacher, hours) VALUES ($1, $2, $3) RETURNING id", subject.Name, subject.Teacher, subject.Hour).Scan(&subject.ID)

		if err != nil {
			panic(err)
		}
		//data := [3]string{subject.Name,subject.Teacher,subject.Hour}
		context.JSON(200, gin.H{"message": "Saved subject", "data": subject})

	} else {
		context.JSON(400, gin.H{"message": "Missing params", "data": ""})
	}
	db.Close()

}

func IndexSubjects(context *gin.Context) {
	db := config.Connect()

	rows, err := db.Query("SELECT * FROM subject")
	if err != nil {
		panic(err)
	}
	subjects := []models.Subject{}

	for rows.Next() {
		subject := models.Subject{}
		err = rows.Scan(&subject.ID, &subject.Name, &subject.Teacher, &subject.Hour)
		if err != nil {
			panic(err)
		}
		subjects = append(subjects, subject)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{"data": subjects})
	db.Close()
}
