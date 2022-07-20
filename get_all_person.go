package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetallPersons(c *gin.Context) {

	Data := []Persons{}

	SQL := "SELECT email,first_name,last_name,date_of_birth,sex from persons"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		rows.Scan(&person.Email, &person.First_name, &person.Last_name, &person.Date_of_birth, &person.Sex)

		Data = append(Data, person)

	}
	res := gin.H{
		"persons": Data,
	}

	c.JSON(http.StatusOK, res)

}
