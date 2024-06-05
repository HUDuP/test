package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

var db *gorm.DB

func (l SqlLogger) Trace(ctx context.Context,begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==============================\n", sql)
}

func main() {
	dsn := "root:P@ssw0rd@tcp(localhost:3306)/meejing?parseTime=true"
	dial := mysql.Open(dsn)

	var err error

	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(Gender{}, Test{}, Customer{})
	//CreateGender("xxxx")
	//GetGender(1)
	//GetGenderByName("Male")
	//UpdateGender2(3, "zzzz")
	// DeleteGender(3)
	// CreateTest(0, "Test1")
	// CreateTest(0, "Test2")
	// CreateTest(0, "Test3")

	// DeleteTest(3)
	// GetTest()
	// db.Migrator().CreateTable(Customer{})
	// CreateCustomer("note", 1)
	GetCustomers()
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name)
	}
}

func GetGenderByName(name string) {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders, "name=?", name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func UpdateGender(id uint,name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

func UpdateGender2(id uint,name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?",id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func CreateTest(code uint, name string) {
	test := Test{Code: code,Name: name}
	db.Create(&test)
}

func GetTest() {
	tests := []Test{}
	db.Find(&tests)
	for _, t := range tests {
		fmt.Printf("%v|%v\n", t.ID,t.Name)
	}
}

func DeleteTest(id uint) {
	db.Unscoped().Delete(&Test{}, id)
}

type Gender struct {
	ID uint
	Name string `gorm:"unique;size(10)"`
}

type Customer struct {
	ID uint
	Name string
	Gender Gender
	GenderID uint
}

type Test struct {
	gorm.Model
	Code uint `gorm:"comment:This is code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

func(t Test) TableName() string {
	return "MyTest"
}