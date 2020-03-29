package main

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestUpdate(t *testing.T) {
	var (
		conn   *gorm.DB
		err    error
		actual = AutoMail{}
	)
	if conn, err = dbConnect(); err != nil {
		t.Fatal(err)
	}
	target := time.Now().Add(time.Second).Truncate(time.Second)
	if err = conn.Model(&AutoMail{}).Where("id=?", 3).Update("send_date_time", target).Error; err != nil {
		t.Fatal(err)
	}
	conn.Model(&AutoMail{}).Where("id=?", 3).Find(&actual)
	if !target.Equal(actual.SendDateTime) {
		t.Errorf("actual %s was not equal target %s", actual.SendDateTime.String(), target.String())
	}
}
func TestRunCommand(t *testing.T) {
	subDateTime := 5 * time.Second
	timeOutDateTime := subDateTime + time.Second
	targetDateTime := time.Now().Add(subDateTime).Truncate(time.Second)
	done := make(chan bool, 1)
	f := func() error {
		done <- true
		return nil
	}
	if err := runCommand(targetDateTime, f); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
		now := time.Now().Truncate(time.Second)
		if !now.Equal(targetDateTime) {
			t.Fatalf("runCommand was not really to run in target time, target %s was not equal to actual %s", targetDateTime, now)
		}
		// t.Log("runCommand is passed")
	case <-time.After(timeOutDateTime):
		t.Fatal("runCommand is time out")
	}
}
