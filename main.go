package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name        string `"json:patient_name"`
	ID          int    `json:"patient_id"`
	Mobile      string `json:"patient_mobile"`
	Message     string `json:"message"`
	IS_SMS_SENT bool   `json:"is_sms_sent"`
}

var patients = []Patient{
	{Name: "Bodhi", ID: 1, Mobile: "+919073423666", Message: "message 1"},
	{Name: "Anonymous", ID: 2, Mobile: "+919836559545", Message: "message 2"},
}

func seed(db *gorm.DB) {

	for _, p := range patients {
		db.Create(&p)
	}
}
func setup(db *gorm.DB) {
	db.AutoMigrate(&Patient{})
	seed(db)
}

func getPatients(c *gin.Context) {
	var NewPatient Patient

	if err := c.BindJSON(&NewPatient); err != nil {
		return
	}

	patients = append(patients, NewPatient)
	c.IndentedJSON(http.StatusCreated, patients)

}

func main() {
	router := gin.Default()
	dsn := "host=localhost user=postgres password=Bo14#08#02 dbname=patients port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't Connect to the Database")
	}
	setup(db)
	var patients []Patient  
	db.Find(&patients)
	for _, p := range patients {
		fmt.Println("Name: ", p.Name, "\nID :", p.ID, "\nMobile :", p.Mobile, "\nMessage :", p.Message)
	}
	router.GET("/patients", getPatients)
	router.Run("localhost:8000")
}