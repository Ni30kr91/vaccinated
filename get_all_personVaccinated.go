package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func personVaccinated(c *gin.Context) {

	Data := []Persons{}

	SQL := "SELECT persons.email,persons.first_name,persons.last_name,persons.date_of_birth,persons.sex from persons, vaccinations WHERE persons.email=vaccinations.recipient"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	person := Persons{}

	for rows.Next() {
		err2 := rows.Scan(&person.Email, &person.First_name, &person.Last_name, &person.Date_of_birth, &person.Sex)
		fmt.Println(err2)
		Data = append(Data, person)

	}
	res := gin.H{
		"person/vaccinated": Data,
	}

	c.JSON(http.StatusOK, res)
	return

}
