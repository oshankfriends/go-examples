package main

import _ "github.com/jinzhu/gorm/dialects/mysql"

import (
	"flag"
	"github.com/jinzhu/gorm"
	"github.com/oshankfriends/go-examples/database/models"
)

var (
	db  *gorm.DB
	dsn = flag.String("DSN", "root:iamplus@tcp(localhost:3306)/company?charset=utf8&parseTime=True", "DSN for MYSQL db")
)

func main() {
	flag.Parse()
	flag.Usage()
	var err error
	if db, err = gorm.Open("mysql", *dsn); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	//	models.CreateTable(db)
	sampleCompany := models.Company{
		Name: "Yahoo",
		Address: models.Address{
			Country:  "USA",
			City:     "Moutain View",
			PostCode: "1600",
		},
		Employees: []models.Employee{
			{
				FirstName:    "John",
				LastName:     "Doe",
				AdhaarNumber: "00-000-0000",
			},
		},
	}
	db.Create(&sampleCompany)
}
