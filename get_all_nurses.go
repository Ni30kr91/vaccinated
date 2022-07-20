package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetallNurses(c *gin.Context) {

	Data := []Nurses{}

	SQL := "SELECT email from nurses"
	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	nurse := Nurses{}

	for rows.Next() {
		rows.Scan(&nurse.Email)

		Data = append(Data, nurse)

	}
	res := gin.H{
		"nurses": Data,
	}
	c.JSON(http.StatusOK, res)
}
