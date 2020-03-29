package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	var (
		err error
	)
	if err = godotenv.Load(); err != nil {
		fmt.Println(err)
	}
}
