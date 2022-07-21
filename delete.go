package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteNursesfromdb(c *gin.Context) {
	reqBody := RemovedNurses{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{

			"Error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, _ := deleteNurses(reqBody)

	if result == false {
		res := gin.H{

			"Error": "Something Went Wrong",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "Successfully Deleted",
		"status":  200,
	}
	c.JSON(http.StatusCreated, res)
}

func deleteNurses(reqbody RemovedNurses) (bool, string) {
	var result = true
	var err_responce = ""

	sqlStatement := `DELETE from nurses WHERE nurses.email=$1 RETURNING email`
	_, err2 := DB.Exec(sqlStatement, reqbody.Email)
	fmt.Println(err2)
	if err2 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}

	sqlStatement2 := `DELETE FROM persons WHERE persons.email=$1`
	_, err3 := DB.Exec(sqlStatement2, reqbody.Email)
	fmt.Println(err3)
	if err3 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	return result, err_responce
}
