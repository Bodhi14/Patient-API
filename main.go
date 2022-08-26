package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Patient struct {
	Name    string `"json:patient_name"`
	ID      int    `json:"patient_id"`
	Mobile  string `json:"patient_mobile"`
	Message string `json:"message"`
}

var patients = []Patient{
	{Name: "A", ID: 1, Mobile: "+91-9073423666", Message: "message"},
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, patients)
}

func main() {
	router := gin.Default()
	router.GET("/patient-details", getData)
	router.Run("localhost:8000")
}
