package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insertnurse(c *gin.Context) {
	reqBody := Persons{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": (err),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	result, result_err := InsertIndb(reqBody)
	if result_err != "" {
		res := gin.H{
			"error":   result_err,
			"success": result,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := gin.H{
		"success": true,
	}
	c.JSON(http.StatusOK, res)

}

func InsertIndb(reqbody Persons) (bool, string) {
	result := true
	err_responce := ""
	sqlStatement := `
	INSERT INTO persons(email , first_name, last_name, date_of_birth, sex)
    VALUES ($1, $2, $3, $4, $5)`
	_, err2 := DB.Exec(sqlStatement, reqbody.Email, reqbody.First_name, reqbody.Last_name, reqbody.Date_of_birth, reqbody.Sex)
	fmt.Println(err2)
	if err2 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}

	sqlStatement2 := `
	INSERT INTO nurses(email)
    VALUES ($1)`
	_, err3 := DB.Exec(sqlStatement2, reqbody.Email)
	fmt.Println(err3)
	if err3 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	return result, err_responce
}
