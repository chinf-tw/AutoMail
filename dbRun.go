package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database
// AutoMail is the database model for automatically send email
type AutoMail struct {
	IsSend       bool `gorm:"default:'false'"`
	TypeInfo     string
	SendFilePath string    `gorm:"not null"`
	SendDateTime time.Time `gorm:"not null"`
	gorm.Model
}

func dbMain() {
	var err error
	if db, err = dbConnect(); err != nil {
		fmt.Println(err)
	}

	if !db.HasTable(&AutoMail{}) {
		if err = db.CreateTable(&AutoMail{}).Error; err != nil {
			fmt.Println(err)
		}
	}
	// count := 10
	// for i := 0; i < count; i++ {
	// 	var newMail = &AutoMail{TypeInfo: "Meeting", SendFilePath: "/data/test.txt"}
	// 	db.Create(&newMail)
	// }

	var a AutoMail
	var qs = []AutoMail{}
	a.IsSend = true
	if err = db.Where("type_info = ?", "Meeting").Find(&qs).Error; err != nil {
		fmt.Println(err)
	}
	for _, q := range qs {
		fmt.Println(q.ID)
	}
}
func dbConnect() (conn *gorm.DB, err error) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Println(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	// fmt.Println(dbURL)
	conn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func runCommand(dateTime time.Time, f func() error) error {
	now := time.Now()
	if now.Before(dateTime) {
		time.Sleep(dateTime.Sub(now))
		if err := f(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("The target dateTime was before the current")
}
