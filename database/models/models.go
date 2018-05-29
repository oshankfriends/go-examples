package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type Company struct {
	ID        int    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string `sql:"size:255;unique;index"`
	Employees []Employee
	Address   Address
}

type Employee struct {
	FirstName    string    `sql:"size:255;index:name_idx"`
	LastName     string    `sql:"size:255;index:name_idx"`
	AdhaarNumber string    `sql:"type:varchar(100);unique" gorm:"column:adhaar"`
	DateOfBirth  time.Time `sql:"DEFAULT:current_timestamp"`
	Address      *Address
	Deleted      bool `sql:"DEFAULT:false"`
}

type Address struct {
	Country  string `gorm:"primary_key"`
	City     string `gorm:"primary_key"`
	PostCode string `gorm:"primary_key"`
	Line1    sql.NullString
	Line2    sql.NullString
}

func CreateTable(db *gorm.DB) {
	db.CreateTable(&Company{})
	db.CreateTable(&Employee{})
	db.CreateTable(&Address{})
}
