package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func nursesVaccinated(c *gin.Context) {

	Data := []Nurses{}

	SQL := "SELECT nurses.email from nurses, vaccinations WHERE nurses.email=vaccinations.recipient"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	nurse := Nurses{}

	for rows.Next() {
		err2 := rows.Scan(&nurse.Email)
		fmt.Println(err2)
		Data = append(Data, nurse)

	}
	res := gin.H{
		"nurses/vaccinated": Data,
	}

	c.JSON(http.StatusOK, res)
	return

}
