package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/boazjr/pullurl/eventpr"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const onPullPath = "/onpull"

//Server accepts incomming PRRequests
//always returns a 200
//TODO: improve security by checking the webhook secret
type Server struct {
	gg                *ggapi
	GithubAccessToken string
	Port              string
}

//Start starts the server
func (s *Server) Start() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	s.gg = &ggapi{gg: github.NewClient(tc)}
	return http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path == onPullPath {
		pr := &eventpr.PREvent{}
		pr.Run(r.Body, s.gg)
	} else {
		log.Println("ERROR: unsupported path", r.URL.Path)
	}
	//TODO: read the docs about reponses
	w.WriteHeader(http.StatusOK)
}
