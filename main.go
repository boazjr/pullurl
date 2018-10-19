package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	fmt.Println("starting")
	s := &server{}
	s.SetupGoGithub()
	s.SetupWebhook()
	err := http.ListenAndServe(":3000", s)
	if err != nil {
		log.Println(err)
	}
}
