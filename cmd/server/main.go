package main

import (
	"github.com/example/play_now/internal/api"
	"log"
)

func main() {
	s := api.NewServer()
	log.Fatal(s.Run())
}
