package server

import (
	"context"

	"github.com/google/go-github/github"
)

type ggapi struct {
	gg *github.Client
}

func (g *ggapi) CreateComment(ctx context.Context, user string, repo string, prN int, comment string) error {
	c := &github.IssueComment{
		Body: &comment,
	}
	_, _, err := g.gg.Issues.CreateComment(context.Background(), user, repo, prN, c)
	return err

}
