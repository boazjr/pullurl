package main

import (
	"log"
	"os"

	"github.com/boazjr/pullurl/server"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("using default PORT")
		port = "3000"
	}
	log.Println("starting server on port:", port)

	gs := os.Getenv("GITHUB_ACCESS_TOKEN")
	if gs == "" {
		log.Printf("ERROR: please set a GITHUB_ACCESS_TOKEN")
		return
	}

	s := &server.Server{
		GithubAccessToken: gs,
		Port:              port,
	}
	log.Println(s.Start())
}
