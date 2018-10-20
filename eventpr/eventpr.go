package eventpr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	"mvdan.cc/xurls"
)

var accptedStatus = []int{http.StatusOK}

type PREvent github.PullRequestEvent

//Run runs a prevent request
func (p *PREvent) Run(body io.ReadCloser, gg githubApi) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return
	}
	defer body.Close()
	pay := &PREvent{}
	err = json.Unmarshal(b, pay)
	if err != nil {
		log.Println(err)
		return
	}
	if pay.Action == nil {
		log.Println("missing Action: ", pay)
		return
	}
	if *pay.Action != "reopened" && *pay.Action != "opened" {
		log.Println("unsupported action: ", *pay.Action)
		return
	}
	err = pay.respond(gg)
	if err != nil {
		log.Println(err)
		return
	}
}

func (p *PREvent) parseBody() (string, error) {
	if p.PullRequest == nil {
		log.Println("Error: pullrequest is empty")
		return "", fmt.Errorf("pullrequest is empty")
	}
	urls := xurls.Relaxed.FindAllString(*p.PullRequest.Body, -1)
	response := ""
	for _, u := range urls {
		response += p.checkURL(u)
	}

	return response, nil
}

func (p *PREvent) respond(gg githubApi) error {
	comment, err := p.parseBody()
	if comment == "" {
		log.Println(err)
		return err
	}
	user, repo, prN, err := p.getGithubInfo()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("creating a comment in: ", user, repo, prN, comment)
	err = gg.CreateComment(context.Background(), user, repo, prN, comment)
	if err != nil {
		return err
	}
	return err
}

type githubApi interface {
	CreateComment(context.Context, string, string, int, string) error
}

func (p *PREvent) checkURL(uin string) string {
	pu, err := url.Parse(uin)
	if err != nil {
		log.Printf("failed to parse url: %s %+v \n", uin, err)
	}
	if pu.Scheme == "" {
		pu.Scheme = "https"
	}
	u := fmt.Sprint(pu)
	resp, err := http.Get(u)
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

// getGithubInfo accepts a PullRequestEvent
// returns username, reponame, pr number, and error
func (p *PREvent) getGithubInfo() (string, string, int, error) {
	parts := strings.Split(p.Repo.GetFullName(), "/")
	if len(parts) != 2 {
		return "", "", 0, fmt.Errorf("failed to get the full name of the repo")
	}
	user, repo := parts[0], parts[1]

	if p.Number == nil {
		return user, repo, 0, fmt.Errorf("pr number is missing %+v", p)
	}
	return user, repo, *p.Number, nil
}
