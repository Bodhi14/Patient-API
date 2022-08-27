package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)

type Patient struct {
	gorm.Model
	Name    string `"json:patient_name"`
	ID      int    `json:"patient_id"`
	Mobile  string `json:"patient_mobile"`
	Message string `json:"message"`
}

func seed(db *gorm.DB) {
	var patients = []Patient{
		{Name: "A", ID: 1, Mobile: "+919073423666", Message: "message"},
	}

	for _, p := range patients {
		db.Create(&p)
	}
}
func setup(db *gorm.DB) {
	db.AutoMigrate(&Patient{})
	seed(db)
}

func main() {
	router := gin.Default()
	db, err := gorm.Open("mysql", "patient.db")
	if err != nil {
		panic("Can't Connect to the Database")
	}
	defer db.Close()
	db.LogMode(true)
	setup(db)
	var patients []Patient
	db.Find(&patients)
	for _, p := range patients {
		fmt.Print("Name: ", p.Name, "ID :", p.ID, "Mobile :", p.Mobile, "Message :", p.Message)
	}
	router.Run("localhost:8000")
}
