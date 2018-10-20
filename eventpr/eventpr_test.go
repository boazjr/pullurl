package eventpr

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func TestPREvent_Run(t *testing.T) {
	type args struct {
		body io.ReadCloser
		gg   githubApi
	}
	mock := &ggapiMock{
		expects: want,
		t:       t,
	}

	tests := []struct {
		name string
		p    *PREvent
		args args
	}{
		{
			name: "test pr event request",
			p:    &PREvent{},
			args: args{
				body: ioutil.NopCloser(bytes.NewReader([]byte(test))),
				gg:   mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PREvent{}
			p.Run(tt.args.body, tt.args.gg)
			if !mock.ok {
				t.Fail()
			}
		})
	}
}

type ggapiMock struct {
	expects string
	t       *testing.T
	ok      bool
	err     error
}

func (g *ggapiMock) CreateComment(ctx context.Context, user string, repo string, prN int, comment string) error {
	if comment == g.expects {
		g.t.Log("passed")
		g.ok = true
		return nil
	}
	g.err = fmt.Errorf("expected %v got %v", g.expects, comment)
	return g.err
}

const want = `200: retrieved url: ` + "`" + `https://www.google.com` + "`" + `
500: failed to retrieve url: ` + "`" + `http://httpstat.us/500` + "`" + `
`

const test = `{
	"action": "reopened",
	"number": 1,
	"pull_request": {
	  "url": "https://api.github.com/repos/boazjr/test/pulls/1",
	  "id": 224343753,
	  "node_id": "MDExOlB1bGxSZXF1ZXN0MjI0MzQzNzUz",
	  "html_url": "https://github.com/boazjr/test/pull/1",
	  "diff_url": "https://github.com/boazjr/test/pull/1.diff",
	  "patch_url": "https://github.com/boazjr/test/pull/1.patch",
	  "issue_url": "https://api.github.com/repos/boazjr/test/issues/1",
	  "number": 1,
	  "state": "closed",
	  "locked": false,
	  "title": "commit name",
	  "user": {
		"login": "boazjr",
		"id": 23318219,
		"node_id": "MDQ6VXNlcjIzMzE4MjE5",
		"avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/boazjr",
		"html_url": "https://github.com/boazjr",
		"followers_url": "https://api.github.com/users/boazjr/followers",
		"following_url": "https://api.github.com/users/boazjr/following{/other_user}",
		"gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
		"organizations_url": "https://api.github.com/users/boazjr/orgs",
		"repos_url": "https://api.github.com/users/boazjr/repos",
		"events_url": "https://api.github.com/users/boazjr/events{/privacy}",
		"received_events_url": "https://api.github.com/users/boazjr/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "body": "www.google.com\r\nAnother line. \n http://httpstat.us/500",
	  "created_at": "2018-10-19T16:30:02Z",
	  "updated_at": "2018-10-19T20:31:00Z",
	  "closed_at": "2018-10-19T20:31:00Z",
	  "merged_at": null,
	  "merge_commit_sha": "076e416f4e43285f661eb8cb9c658e3cba3fe3f5",
	  "assignee": null,
	  "assignees": [
  
	  ],
	  "requested_reviewers": [
  
	  ],
	  "requested_teams": [
  
	  ],
	  "labels": [
  
	  ],
	  "milestone": null,
	  "commits_url": "https://api.github.com/repos/boazjr/test/pulls/1/commits",
	  "review_comments_url": "https://api.github.com/repos/boazjr/test/pulls/1/comments",
	  "review_comment_url": "https://api.github.com/repos/boazjr/test/pulls/comments{/number}",
	  "comments_url": "https://api.github.com/repos/boazjr/test/issues/1/comments",
	  "statuses_url": "https://api.github.com/repos/boazjr/test/statuses/0ea11caabb0743e73df84b851b119ef1174f2a28",
	  "head": {
		"label": "boazjr:test1",
		"ref": "test1",
		"sha": "0ea11caabb0743e73df84b851b119ef1174f2a28",
		"user": {
		  "login": "boazjr",
		  "id": 23318219,
		  "node_id": "MDQ6VXNlcjIzMzE4MjE5",
		  "avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/boazjr",
		  "html_url": "https://github.com/boazjr",
		  "followers_url": "https://api.github.com/users/boazjr/followers",
		  "following_url": "https://api.github.com/users/boazjr/following{/other_user}",
		  "gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
		  "organizations_url": "https://api.github.com/users/boazjr/orgs",
		  "repos_url": "https://api.github.com/users/boazjr/repos",
		  "events_url": "https://api.github.com/users/boazjr/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/boazjr/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 153809803,
		  "node_id": "MDEwOlJlcG9zaXRvcnkxNTM4MDk4MDM=",
		  "name": "test",
		  "full_name": "boazjr/test",
		  "private": false,
		  "owner": {
			"login": "boazjr",
			"id": 23318219,
			"node_id": "MDQ6VXNlcjIzMzE4MjE5",
			"avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/boazjr",
			"html_url": "https://github.com/boazjr",
			"followers_url": "https://api.github.com/users/boazjr/followers",
			"following_url": "https://api.github.com/users/boazjr/following{/other_user}",
			"gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
			"organizations_url": "https://api.github.com/users/boazjr/orgs",
			"repos_url": "https://api.github.com/users/boazjr/repos",
			"events_url": "https://api.github.com/users/boazjr/events{/privacy}",
			"received_events_url": "https://api.github.com/users/boazjr/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "html_url": "https://github.com/boazjr/test",
		  "description": null,
		  "fork": false,
		  "url": "https://api.github.com/repos/boazjr/test",
		  "forks_url": "https://api.github.com/repos/boazjr/test/forks",
		  "keys_url": "https://api.github.com/repos/boazjr/test/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/boazjr/test/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/boazjr/test/teams",
		  "hooks_url": "https://api.github.com/repos/boazjr/test/hooks",
		  "issue_events_url": "https://api.github.com/repos/boazjr/test/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/boazjr/test/events",
		  "assignees_url": "https://api.github.com/repos/boazjr/test/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/boazjr/test/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/boazjr/test/tags",
		  "blobs_url": "https://api.github.com/repos/boazjr/test/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/boazjr/test/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/boazjr/test/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/boazjr/test/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/boazjr/test/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/boazjr/test/languages",
		  "stargazers_url": "https://api.github.com/repos/boazjr/test/stargazers",
		  "contributors_url": "https://api.github.com/repos/boazjr/test/contributors",
		  "subscribers_url": "https://api.github.com/repos/boazjr/test/subscribers",
		  "subscription_url": "https://api.github.com/repos/boazjr/test/subscription",
		  "commits_url": "https://api.github.com/repos/boazjr/test/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/boazjr/test/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/boazjr/test/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/boazjr/test/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/boazjr/test/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/boazjr/test/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/boazjr/test/merges",
		  "archive_url": "https://api.github.com/repos/boazjr/test/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/boazjr/test/downloads",
		  "issues_url": "https://api.github.com/repos/boazjr/test/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/boazjr/test/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/boazjr/test/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/boazjr/test/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/boazjr/test/labels{/name}",
		  "releases_url": "https://api.github.com/repos/boazjr/test/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/boazjr/test/deployments",
		  "created_at": "2018-10-19T16:12:09Z",
		  "updated_at": "2018-10-19T16:27:46Z",
		  "pushed_at": "2018-10-19T16:55:25Z",
		  "git_url": "git://github.com/boazjr/test.git",
		  "ssh_url": "git@github.com:boazjr/test.git",
		  "clone_url": "https://github.com/boazjr/test.git",
		  "svn_url": "https://github.com/boazjr/test",
		  "homepage": null,
		  "size": 1,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": true,
		  "has_pages": false,
		  "forks_count": 0,
		  "mirror_url": null,
		  "archived": false,
		  "open_issues_count": 1,
		  "license": null,
		  "forks": 0,
		  "open_issues": 1,
		  "watchers": 0,
		  "default_branch": "master"
		}
	  },
	  "base": {
		"label": "boazjr:master",
		"ref": "master",
		"sha": "29241f5880f6937519ca133ad9f9f99afd086e5e",
		"user": {
		  "login": "boazjr",
		  "id": 23318219,
		  "node_id": "MDQ6VXNlcjIzMzE4MjE5",
		  "avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/boazjr",
		  "html_url": "https://github.com/boazjr",
		  "followers_url": "https://api.github.com/users/boazjr/followers",
		  "following_url": "https://api.github.com/users/boazjr/following{/other_user}",
		  "gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
		  "organizations_url": "https://api.github.com/users/boazjr/orgs",
		  "repos_url": "https://api.github.com/users/boazjr/repos",
		  "events_url": "https://api.github.com/users/boazjr/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/boazjr/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 153809803,
		  "node_id": "MDEwOlJlcG9zaXRvcnkxNTM4MDk4MDM=",
		  "name": "test",
		  "full_name": "boazjr/test",
		  "private": false,
		  "owner": {
			"login": "boazjr",
			"id": 23318219,
			"node_id": "MDQ6VXNlcjIzMzE4MjE5",
			"avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/boazjr",
			"html_url": "https://github.com/boazjr",
			"followers_url": "https://api.github.com/users/boazjr/followers",
			"following_url": "https://api.github.com/users/boazjr/following{/other_user}",
			"gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
			"organizations_url": "https://api.github.com/users/boazjr/orgs",
			"repos_url": "https://api.github.com/users/boazjr/repos",
			"events_url": "https://api.github.com/users/boazjr/events{/privacy}",
			"received_events_url": "https://api.github.com/users/boazjr/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "html_url": "https://github.com/boazjr/test",
		  "description": null,
		  "fork": false,
		  "url": "https://api.github.com/repos/boazjr/test",
		  "forks_url": "https://api.github.com/repos/boazjr/test/forks",
		  "keys_url": "https://api.github.com/repos/boazjr/test/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/boazjr/test/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/boazjr/test/teams",
		  "hooks_url": "https://api.github.com/repos/boazjr/test/hooks",
		  "issue_events_url": "https://api.github.com/repos/boazjr/test/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/boazjr/test/events",
		  "assignees_url": "https://api.github.com/repos/boazjr/test/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/boazjr/test/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/boazjr/test/tags",
		  "blobs_url": "https://api.github.com/repos/boazjr/test/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/boazjr/test/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/boazjr/test/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/boazjr/test/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/boazjr/test/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/boazjr/test/languages",
		  "stargazers_url": "https://api.github.com/repos/boazjr/test/stargazers",
		  "contributors_url": "https://api.github.com/repos/boazjr/test/contributors",
		  "subscribers_url": "https://api.github.com/repos/boazjr/test/subscribers",
		  "subscription_url": "https://api.github.com/repos/boazjr/test/subscription",
		  "commits_url": "https://api.github.com/repos/boazjr/test/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/boazjr/test/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/boazjr/test/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/boazjr/test/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/boazjr/test/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/boazjr/test/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/boazjr/test/merges",
		  "archive_url": "https://api.github.com/repos/boazjr/test/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/boazjr/test/downloads",
		  "issues_url": "https://api.github.com/repos/boazjr/test/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/boazjr/test/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/boazjr/test/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/boazjr/test/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/boazjr/test/labels{/name}",
		  "releases_url": "https://api.github.com/repos/boazjr/test/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/boazjr/test/deployments",
		  "created_at": "2018-10-19T16:12:09Z",
		  "updated_at": "2018-10-19T16:27:46Z",
		  "pushed_at": "2018-10-19T16:55:25Z",
		  "git_url": "git://github.com/boazjr/test.git",
		  "ssh_url": "git@github.com:boazjr/test.git",
		  "clone_url": "https://github.com/boazjr/test.git",
		  "svn_url": "https://github.com/boazjr/test",
		  "homepage": null,
		  "size": 1,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": true,
		  "has_pages": false,
		  "forks_count": 0,
		  "mirror_url": null,
		  "archived": false,
		  "open_issues_count": 1,
		  "license": null,
		  "forks": 0,
		  "open_issues": 1,
		  "watchers": 0,
		  "default_branch": "master"
		}
	  },
	  "_links": {
		"self": {
		  "href": "https://api.github.com/repos/boazjr/test/pulls/1"
		},
		"html": {
		  "href": "https://github.com/boazjr/test/pull/1"
		},
		"issue": {
		  "href": "https://api.github.com/repos/boazjr/test/issues/1"
		},
		"comments": {
		  "href": "https://api.github.com/repos/boazjr/test/issues/1/comments"
		},
		"review_comments": {
		  "href": "https://api.github.com/repos/boazjr/test/pulls/1/comments"
		},
		"review_comment": {
		  "href": "https://api.github.com/repos/boazjr/test/pulls/comments{/number}"
		},
		"commits": {
		  "href": "https://api.github.com/repos/boazjr/test/pulls/1/commits"
		},
		"statuses": {
		  "href": "https://api.github.com/repos/boazjr/test/statuses/0ea11caabb0743e73df84b851b119ef1174f2a28"
		}
	  },
	  "author_association": "OWNER",
	  "merged": false,
	  "mergeable": true,
	  "rebaseable": true,
	  "mergeable_state": "clean",
	  "merged_by": null,
	  "comments": 0,
	  "review_comments": 0,
	  "maintainer_can_modify": false,
	  "commits": 1,
	  "additions": 2,
	  "deletions": 0,
	  "changed_files": 1
	},
	"repository": {
	  "id": 153809803,
	  "node_id": "MDEwOlJlcG9zaXRvcnkxNTM4MDk4MDM=",
	  "name": "test",
	  "full_name": "boazjr/test",
	  "private": false,
	  "owner": {
		"login": "boazjr",
		"id": 23318219,
		"node_id": "MDQ6VXNlcjIzMzE4MjE5",
		"avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/boazjr",
		"html_url": "https://github.com/boazjr",
		"followers_url": "https://api.github.com/users/boazjr/followers",
		"following_url": "https://api.github.com/users/boazjr/following{/other_user}",
		"gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
		"organizations_url": "https://api.github.com/users/boazjr/orgs",
		"repos_url": "https://api.github.com/users/boazjr/repos",
		"events_url": "https://api.github.com/users/boazjr/events{/privacy}",
		"received_events_url": "https://api.github.com/users/boazjr/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "html_url": "https://github.com/boazjr/test",
	  "description": null,
	  "fork": false,
	  "url": "https://api.github.com/repos/boazjr/test",
	  "forks_url": "https://api.github.com/repos/boazjr/test/forks",
	  "keys_url": "https://api.github.com/repos/boazjr/test/keys{/key_id}",
	  "collaborators_url": "https://api.github.com/repos/boazjr/test/collaborators{/collaborator}",
	  "teams_url": "https://api.github.com/repos/boazjr/test/teams",
	  "hooks_url": "https://api.github.com/repos/boazjr/test/hooks",
	  "issue_events_url": "https://api.github.com/repos/boazjr/test/issues/events{/number}",
	  "events_url": "https://api.github.com/repos/boazjr/test/events",
	  "assignees_url": "https://api.github.com/repos/boazjr/test/assignees{/user}",
	  "branches_url": "https://api.github.com/repos/boazjr/test/branches{/branch}",
	  "tags_url": "https://api.github.com/repos/boazjr/test/tags",
	  "blobs_url": "https://api.github.com/repos/boazjr/test/git/blobs{/sha}",
	  "git_tags_url": "https://api.github.com/repos/boazjr/test/git/tags{/sha}",
	  "git_refs_url": "https://api.github.com/repos/boazjr/test/git/refs{/sha}",
	  "trees_url": "https://api.github.com/repos/boazjr/test/git/trees{/sha}",
	  "statuses_url": "https://api.github.com/repos/boazjr/test/statuses/{sha}",
	  "languages_url": "https://api.github.com/repos/boazjr/test/languages",
	  "stargazers_url": "https://api.github.com/repos/boazjr/test/stargazers",
	  "contributors_url": "https://api.github.com/repos/boazjr/test/contributors",
	  "subscribers_url": "https://api.github.com/repos/boazjr/test/subscribers",
	  "subscription_url": "https://api.github.com/repos/boazjr/test/subscription",
	  "commits_url": "https://api.github.com/repos/boazjr/test/commits{/sha}",
	  "git_commits_url": "https://api.github.com/repos/boazjr/test/git/commits{/sha}",
	  "comments_url": "https://api.github.com/repos/boazjr/test/comments{/number}",
	  "issue_comment_url": "https://api.github.com/repos/boazjr/test/issues/comments{/number}",
	  "contents_url": "https://api.github.com/repos/boazjr/test/contents/{+path}",
	  "compare_url": "https://api.github.com/repos/boazjr/test/compare/{base}...{head}",
	  "merges_url": "https://api.github.com/repos/boazjr/test/merges",
	  "archive_url": "https://api.github.com/repos/boazjr/test/{archive_format}{/ref}",
	  "downloads_url": "https://api.github.com/repos/boazjr/test/downloads",
	  "issues_url": "https://api.github.com/repos/boazjr/test/issues{/number}",
	  "pulls_url": "https://api.github.com/repos/boazjr/test/pulls{/number}",
	  "milestones_url": "https://api.github.com/repos/boazjr/test/milestones{/number}",
	  "notifications_url": "https://api.github.com/repos/boazjr/test/notifications{?since,all,participating}",
	  "labels_url": "https://api.github.com/repos/boazjr/test/labels{/name}",
	  "releases_url": "https://api.github.com/repos/boazjr/test/releases{/id}",
	  "deployments_url": "https://api.github.com/repos/boazjr/test/deployments",
	  "created_at": "2018-10-19T16:12:09Z",
	  "updated_at": "2018-10-19T16:27:46Z",
	  "pushed_at": "2018-10-19T16:55:25Z",
	  "git_url": "git://github.com/boazjr/test.git",
	  "ssh_url": "git@github.com:boazjr/test.git",
	  "clone_url": "https://github.com/boazjr/test.git",
	  "svn_url": "https://github.com/boazjr/test",
	  "homepage": null,
	  "size": 1,
	  "stargazers_count": 0,
	  "watchers_count": 0,
	  "language": null,
	  "has_issues": true,
	  "has_projects": true,
	  "has_downloads": true,
	  "has_wiki": true,
	  "has_pages": false,
	  "forks_count": 0,
	  "mirror_url": null,
	  "archived": false,
	  "open_issues_count": 1,
	  "license": null,
	  "forks": 0,
	  "open_issues": 1,
	  "watchers": 0,
	  "default_branch": "master"
	},
	"sender": {
	  "login": "boazjr",
	  "id": 23318219,
	  "node_id": "MDQ6VXNlcjIzMzE4MjE5",
	  "avatar_url": "https://avatars2.githubusercontent.com/u/23318219?v=4",
	  "gravatar_id": "",
	  "url": "https://api.github.com/users/boazjr",
	  "html_url": "https://github.com/boazjr",
	  "followers_url": "https://api.github.com/users/boazjr/followers",
	  "following_url": "https://api.github.com/users/boazjr/following{/other_user}",
	  "gists_url": "https://api.github.com/users/boazjr/gists{/gist_id}",
	  "starred_url": "https://api.github.com/users/boazjr/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/boazjr/subscriptions",
	  "organizations_url": "https://api.github.com/users/boazjr/orgs",
	  "repos_url": "https://api.github.com/users/boazjr/repos",
	  "events_url": "https://api.github.com/users/boazjr/events{/privacy}",
	  "received_events_url": "https://api.github.com/users/boazjr/received_events",
	  "type": "User",
	  "site_admin": false
	}
  }
`
