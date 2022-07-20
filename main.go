package main

import (

	//"encoding/csv"

	"github.com/gin-gonic/gin"
)

type Nurses struct {
	Email string `json:"email"`
	// First_name    string `json:"first_name"`
	// Last_name     string `json:"last_name"`
	// Date_of_birth string `json:"date_of_birth"`
	// Sex           string `json:"sex"`
}

type Cities struct {
	City  string `json:"city"`
	State string `json:"state"`
}
type Persons struct {
	Email         string `json:"email" binding:"email"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Date_of_birth string `json:"date_of_birth"`
	Sex           string `json:"sex"`
}

type Vaccination_sites struct {
	Site    string `json:"site"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Open    string `json:"open"`
	Close   string `json:"close"`
}

type Vaccination struct {
	Recipient        string `json:"recipient"`
	Vaccination_time string `json:"vaccination_time "`
	Vaccine          string `json:"vaccine"`
	Site             string `json:"site"`
	Nurse            string `json:"nurse"`
	Comments         string `json:"comments"`
}

type Vaccine_providers struct {
	Company string `json:"company"`
}

func main() {
	createDBConnection()

	//importCSV()

	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {

	r.GET("nurses", GetallNurses)
	r.GET("persons", GetallPersons)
	r.GET("person/vaccinated", personVaccinated)
	r.GET("nurses/vaccinated", nursesVaccinated)

	// Get all nurses
	// Get all persons
	// Get all person who are vaccinated
	// Get all nursesvaccinated

}
