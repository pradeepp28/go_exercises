package main

import (
	"go_exercises/playquiz/game"
	"log"
)

func main() {
	qz, err := game.NewQuiz()
	if err != nil {
		log.Fatal(err)
	}

	_ = qz
	// TODO
}
