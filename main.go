package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Patient struct {
	gorm.Model
	Name    string `"json:patient_name"`
	ID      int    `json:"patient_id"`
	Mobile  string `json:"patient_mobile"`
	Message string `json:"message"`
}

var patients = []Patient{
	{Name: "A", ID: 1, Mobile: "+919073423666", Message: "message"},
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, patients)
}

func setup(db *gorm.db) {
}

func main() {
	router := gin.Default()
	router.GET("/patient-details", getData)
	router.Run("localhost:8000")
}
