package main

import (
	"github.com/bonNope/introTask/internal/pkg/app"
	"log"
)

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
