package main

import (
	"log"

	"github.com/mayusGomez/testGolangCoverage/sender"
)

func main() {
	media := sender.NewEmail()
	err := media.Send("Hello World")
	if err != nil {
		log.Println("error:", err)
	}
}
