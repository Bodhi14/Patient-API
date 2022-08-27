package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't Connect to the Database")
	}
	setup(db)
	var patients []Patient
	db.Find(&patients)
	for _, p := range patients {
		fmt.Print("Name: ", p.Name, "ID :", p.ID, "Mobile :", p.Mobile, "Message :", p.Message)
	}
	router.Run("localhost:8000")
}
