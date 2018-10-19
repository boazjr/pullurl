package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	gogithub "github.com/google/go-github/github"
	"gopkg.in/go-playground/webhooks.v5/github"
	"mvdan.cc/xurls"
)

const onPullPath = "/onpull"

var accptedStatus = []int{http.StatusOK}

type server struct {
	gg   *gogithub.Client
	hook *github.Webhook
	get  Getter
}

func (s *server) SetupGoGithub() {
	s.gg = gogithub.NewClient(nil)
}

func (s *server) SetupWebhook() {
	gs := os.Getenv("GITHUB_SECRET")
	if gs == "" {
		log.Fatal("GITHUB_SECRET envar is not set")
	}

	hook, err := github.New(github.Options.Secret(gs))
	if err != nil {
		log.Fatalf("failed to create github hook client %+v", err)
	}
	s.hook = hook
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == onPullPath {
		s.handlePull(w, r)
	} else {
		log.Println("ERROR: unsupported path", r.URL.Path)
	}
	// in both cases we return ok
	//TODO: read the docs about reponses
	w.WriteHeader(http.StatusOK)
	return
}

func (s *server) handlePull(w http.ResponseWriter, r *http.Request) {
	payload, err := s.hook.Parse(r, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
			log.Printf("event not requested: %+v", err)
		}
	}
	switch payload.(type) {
	case github.PullRequestPayload:
		pullRequest := payload.(github.PullRequestPayload)
		if err != nil {
			log.Println(err)
		}
		// Do whatever you want from here...
		fmt.Printf("%+v \n", pullRequest)
		res, pri := s.parseBody(pullRequest)
		if res == "" {
			return
		}
		err = s.respond(res, pri)
		if err != nil {
			log.Println(err)
		}

	default:
		log.Println("unsupported payload")
	}
}

func (s *server) respond(payload string, pri int64) error {
	ctx := context.Background()
	comment := &gogithub.PullRequestComment{
		Body: &payload,
	}
	//TODO: pri number type is inconsistent with both libraries.
	_, _, err := s.gg.PullRequests.CreateComment(ctx, "boazjr", "test", int(pri), comment)
	return err
}

func (s *server) parseBody(p github.PullRequestPayload) (string, int64) {
	urls := xurls.Relaxed().FindAllString(p.PullRequest.Body, -1)
	response := ""
	for _, u := range urls {
		response += s.CheckURL(u)
	}
	return response, p.Number
}

func (s *server) CheckURL(uin string) string {
	pu, err := url.Parse(uin)
	if err != nil {
		log.Printf("failed to parse url: %s %+v \n", uin, err)
	}
	if pu.Scheme == "" {
		pu.Scheme = "https"
	}
	u := fmt.Sprint(pu)
	resp, err := s.get.Get(u)
	if err != nil {
		log.Printf("failed to retrieve url: %s %+v\n", u, err)
		return fmt.Sprintf("Error: failed to retrieve url: %+s, %+v\n", u, err)
	}
	for _, ac := range accptedStatus {
		if resp.StatusCode == ac {
			return fmt.Sprintf("%d: retrieved url: `%s`\n", resp.StatusCode, u)
		}
	}
	return fmt.Sprintf("%d: failed to retrieve url: `%s`\n", resp.StatusCode, u)
}

//Getter allows mocking the golang http client to make get requests
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}
