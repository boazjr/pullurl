package main

import (
	"encoding/json"
	"net/http"
	"testing"

	gogithub "github.com/google/go-github/github"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func Test_server_parseBody(t *testing.T) {
	type fields struct {
		gg   *gogithub.Client
		hook *github.Webhook
		get  Getter
	}
	type args struct {
		p github.PullRequestPayload
	}
	p := &github.PullRequestPayload{}
	json.Unmarshal([]byte(test), p)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "",
			fields: fields{
				get: http.DefaultClient,
			},
			args: args{
				p: *p,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				gg:   tt.fields.gg,
				hook: tt.fields.hook,
				get:  tt.fields.get,
			}
			if got := s.parseBody(tt.args.p); got != tt.want {
				t.Errorf("server.parseBody() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

const want = `200: retrieved url: ` + "`" + `https://www.google.com` + "`" + `
500: failed to retrieve url: ` + "`" + `http://httpstat.us/500` + "`" + `
`

const test = `{
	"action": "closed",
	"number": 1,
	"pull_request": {
	  "url": "https://api.github.com/repos/boazjr/test/pulls/1",
	  "id": 191568743,
	  "node_id": "MDExOlB1bGxSZXF1ZXN0MTkxNTY4NzQz",
	  "html_url": "https://github.com/Codertocat/Hello-World/pull/1",
	  "diff_url": "https://github.com/Codertocat/Hello-World/pull/1.diff",
	  "patch_url": "https://github.com/Codertocat/Hello-World/pull/1.patch",
	  "issue_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/1",
	  "number": 1,
	  "state": "closed",
	  "locked": false,
	  "title": "Update the README with new information",
	  "user": {
		"login": "Codertocat",
		"id": 21031067,
		"node_id": "MDQ6VXNlcjIxMDMxMDY3",
		"avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/Codertocat",
		"html_url": "https://github.com/Codertocat",
		"followers_url": "https://api.github.com/users/Codertocat/followers",
		"following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
		"gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
		"organizations_url": "https://api.github.com/users/Codertocat/orgs",
		"repos_url": "https://api.github.com/users/Codertocat/repos",
		"events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Codertocat/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "body": "www.google.com    http://httpstat.us/500 some other string",
	  "created_at": "2018-05-30T20:18:30Z",
	  "updated_at": "2018-05-30T20:18:50Z",
	  "closed_at": "2018-05-30T20:18:50Z",
	  "merged_at": null,
	  "merge_commit_sha": "414cb0069601a32b00bd122a2380cd283626a8e5",
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
	  "commits_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls/1/commits",
	  "review_comments_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls/1/comments",
	  "review_comment_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments{/number}",
	  "comments_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/1/comments",
	  "statuses_url": "https://api.github.com/repos/Codertocat/Hello-World/statuses/34c5c7793cb3b279e22454cb6750c80560547b3a",
	  "head": {
		"label": "Codertocat:changes",
		"ref": "changes",
		"sha": "34c5c7793cb3b279e22454cb6750c80560547b3a",
		"user": {
		  "login": "Codertocat",
		  "id": 21031067,
		  "node_id": "MDQ6VXNlcjIxMDMxMDY3",
		  "avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/Codertocat",
		  "html_url": "https://github.com/Codertocat",
		  "followers_url": "https://api.github.com/users/Codertocat/followers",
		  "following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
		  "gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
		  "organizations_url": "https://api.github.com/users/Codertocat/orgs",
		  "repos_url": "https://api.github.com/users/Codertocat/repos",
		  "events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/Codertocat/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 135493233,
		  "node_id": "MDEwOlJlcG9zaXRvcnkxMzU0OTMyMzM=",
		  "name": "Hello-World",
		  "full_name": "Codertocat/Hello-World",
		  "owner": {
			"login": "Codertocat",
			"id": 21031067,
			"node_id": "MDQ6VXNlcjIxMDMxMDY3",
			"avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/Codertocat",
			"html_url": "https://github.com/Codertocat",
			"followers_url": "https://api.github.com/users/Codertocat/followers",
			"following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
			"gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
			"organizations_url": "https://api.github.com/users/Codertocat/orgs",
			"repos_url": "https://api.github.com/users/Codertocat/repos",
			"events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/Codertocat/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "private": false,
		  "html_url": "https://github.com/Codertocat/Hello-World",
		  "description": null,
		  "fork": false,
		  "url": "https://api.github.com/repos/Codertocat/Hello-World",
		  "forks_url": "https://api.github.com/repos/Codertocat/Hello-World/forks",
		  "keys_url": "https://api.github.com/repos/Codertocat/Hello-World/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/Codertocat/Hello-World/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/Codertocat/Hello-World/teams",
		  "hooks_url": "https://api.github.com/repos/Codertocat/Hello-World/hooks",
		  "issue_events_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/Codertocat/Hello-World/events",
		  "assignees_url": "https://api.github.com/repos/Codertocat/Hello-World/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/Codertocat/Hello-World/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/Codertocat/Hello-World/tags",
		  "blobs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/Codertocat/Hello-World/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/Codertocat/Hello-World/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/Codertocat/Hello-World/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/Codertocat/Hello-World/languages",
		  "stargazers_url": "https://api.github.com/repos/Codertocat/Hello-World/stargazers",
		  "contributors_url": "https://api.github.com/repos/Codertocat/Hello-World/contributors",
		  "subscribers_url": "https://api.github.com/repos/Codertocat/Hello-World/subscribers",
		  "subscription_url": "https://api.github.com/repos/Codertocat/Hello-World/subscription",
		  "commits_url": "https://api.github.com/repos/Codertocat/Hello-World/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/Codertocat/Hello-World/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/Codertocat/Hello-World/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/Codertocat/Hello-World/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/Codertocat/Hello-World/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/Codertocat/Hello-World/merges",
		  "archive_url": "https://api.github.com/repos/Codertocat/Hello-World/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/Codertocat/Hello-World/downloads",
		  "issues_url": "https://api.github.com/repos/Codertocat/Hello-World/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/Codertocat/Hello-World/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/Codertocat/Hello-World/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/Codertocat/Hello-World/labels{/name}",
		  "releases_url": "https://api.github.com/repos/Codertocat/Hello-World/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/Codertocat/Hello-World/deployments",
		  "created_at": "2018-05-30T20:18:04Z",
		  "updated_at": "2018-05-30T20:18:50Z",
		  "pushed_at": "2018-05-30T20:18:48Z",
		  "git_url": "git://github.com/Codertocat/Hello-World.git",
		  "ssh_url": "git@github.com:Codertocat/Hello-World.git",
		  "clone_url": "https://github.com/Codertocat/Hello-World.git",
		  "svn_url": "https://github.com/Codertocat/Hello-World",
		  "homepage": null,
		  "size": 0,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": true,
		  "has_pages": true,
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
		"label": "Codertocat:master",
		"ref": "master",
		"sha": "a10867b14bb761a232cd80139fbd4c0d33264240",
		"user": {
		  "login": "Codertocat",
		  "id": 21031067,
		  "node_id": "MDQ6VXNlcjIxMDMxMDY3",
		  "avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/Codertocat",
		  "html_url": "https://github.com/Codertocat",
		  "followers_url": "https://api.github.com/users/Codertocat/followers",
		  "following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
		  "gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
		  "organizations_url": "https://api.github.com/users/Codertocat/orgs",
		  "repos_url": "https://api.github.com/users/Codertocat/repos",
		  "events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/Codertocat/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 135493233,
		  "node_id": "MDEwOlJlcG9zaXRvcnkxMzU0OTMyMzM=",
		  "name": "Hello-World",
		  "full_name": "Codertocat/Hello-World",
		  "owner": {
			"login": "Codertocat",
			"id": 21031067,
			"node_id": "MDQ6VXNlcjIxMDMxMDY3",
			"avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/Codertocat",
			"html_url": "https://github.com/Codertocat",
			"followers_url": "https://api.github.com/users/Codertocat/followers",
			"following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
			"gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
			"organizations_url": "https://api.github.com/users/Codertocat/orgs",
			"repos_url": "https://api.github.com/users/Codertocat/repos",
			"events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/Codertocat/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "private": false,
		  "html_url": "https://github.com/Codertocat/Hello-World",
		  "description": null,
		  "fork": false,
		  "url": "https://api.github.com/repos/Codertocat/Hello-World",
		  "forks_url": "https://api.github.com/repos/Codertocat/Hello-World/forks",
		  "keys_url": "https://api.github.com/repos/Codertocat/Hello-World/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/Codertocat/Hello-World/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/Codertocat/Hello-World/teams",
		  "hooks_url": "https://api.github.com/repos/Codertocat/Hello-World/hooks",
		  "issue_events_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/Codertocat/Hello-World/events",
		  "assignees_url": "https://api.github.com/repos/Codertocat/Hello-World/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/Codertocat/Hello-World/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/Codertocat/Hello-World/tags",
		  "blobs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/Codertocat/Hello-World/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/Codertocat/Hello-World/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/Codertocat/Hello-World/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/Codertocat/Hello-World/languages",
		  "stargazers_url": "https://api.github.com/repos/Codertocat/Hello-World/stargazers",
		  "contributors_url": "https://api.github.com/repos/Codertocat/Hello-World/contributors",
		  "subscribers_url": "https://api.github.com/repos/Codertocat/Hello-World/subscribers",
		  "subscription_url": "https://api.github.com/repos/Codertocat/Hello-World/subscription",
		  "commits_url": "https://api.github.com/repos/Codertocat/Hello-World/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/Codertocat/Hello-World/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/Codertocat/Hello-World/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/Codertocat/Hello-World/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/Codertocat/Hello-World/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/Codertocat/Hello-World/merges",
		  "archive_url": "https://api.github.com/repos/Codertocat/Hello-World/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/Codertocat/Hello-World/downloads",
		  "issues_url": "https://api.github.com/repos/Codertocat/Hello-World/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/Codertocat/Hello-World/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/Codertocat/Hello-World/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/Codertocat/Hello-World/labels{/name}",
		  "releases_url": "https://api.github.com/repos/Codertocat/Hello-World/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/Codertocat/Hello-World/deployments",
		  "created_at": "2018-05-30T20:18:04Z",
		  "updated_at": "2018-05-30T20:18:50Z",
		  "pushed_at": "2018-05-30T20:18:48Z",
		  "git_url": "git://github.com/Codertocat/Hello-World.git",
		  "ssh_url": "git@github.com:Codertocat/Hello-World.git",
		  "clone_url": "https://github.com/Codertocat/Hello-World.git",
		  "svn_url": "https://github.com/Codertocat/Hello-World",
		  "homepage": null,
		  "size": 0,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": true,
		  "has_pages": true,
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
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/1"
		},
		"html": {
		  "href": "https://github.com/Codertocat/Hello-World/pull/1"
		},
		"issue": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/issues/1"
		},
		"comments": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/issues/1/comments"
		},
		"review_comments": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/1/comments"
		},
		"review_comment": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments{/number}"
		},
		"commits": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/1/commits"
		},
		"statuses": {
		  "href": "https://api.github.com/repos/Codertocat/Hello-World/statuses/34c5c7793cb3b279e22454cb6750c80560547b3a"
		}
	  },
	  "author_association": "OWNER",
	  "merged": false,
	  "mergeable": true,
	  "rebaseable": true,
	  "mergeable_state": "clean",
	  "merged_by": null,
	  "comments": 0,
	  "review_comments": 1,
	  "maintainer_can_modify": false,
	  "commits": 1,
	  "additions": 1,
	  "deletions": 1,
	  "changed_files": 1
	},
	"repository": {
	  "id": 135493233,
	  "node_id": "MDEwOlJlcG9zaXRvcnkxMzU0OTMyMzM=",
	  "name": "Hello-World",
	  "full_name": "Codertocat/Hello-World",
	  "owner": {
		"login": "Codertocat",
		"id": 21031067,
		"node_id": "MDQ6VXNlcjIxMDMxMDY3",
		"avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/Codertocat",
		"html_url": "https://github.com/Codertocat",
		"followers_url": "https://api.github.com/users/Codertocat/followers",
		"following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
		"gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
		"organizations_url": "https://api.github.com/users/Codertocat/orgs",
		"repos_url": "https://api.github.com/users/Codertocat/repos",
		"events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Codertocat/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "private": false,
	  "html_url": "https://github.com/Codertocat/Hello-World",
	  "description": null,
	  "fork": false,
	  "url": "https://api.github.com/repos/Codertocat/Hello-World",
	  "forks_url": "https://api.github.com/repos/Codertocat/Hello-World/forks",
	  "keys_url": "https://api.github.com/repos/Codertocat/Hello-World/keys{/key_id}",
	  "collaborators_url": "https://api.github.com/repos/Codertocat/Hello-World/collaborators{/collaborator}",
	  "teams_url": "https://api.github.com/repos/Codertocat/Hello-World/teams",
	  "hooks_url": "https://api.github.com/repos/Codertocat/Hello-World/hooks",
	  "issue_events_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/events{/number}",
	  "events_url": "https://api.github.com/repos/Codertocat/Hello-World/events",
	  "assignees_url": "https://api.github.com/repos/Codertocat/Hello-World/assignees{/user}",
	  "branches_url": "https://api.github.com/repos/Codertocat/Hello-World/branches{/branch}",
	  "tags_url": "https://api.github.com/repos/Codertocat/Hello-World/tags",
	  "blobs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/blobs{/sha}",
	  "git_tags_url": "https://api.github.com/repos/Codertocat/Hello-World/git/tags{/sha}",
	  "git_refs_url": "https://api.github.com/repos/Codertocat/Hello-World/git/refs{/sha}",
	  "trees_url": "https://api.github.com/repos/Codertocat/Hello-World/git/trees{/sha}",
	  "statuses_url": "https://api.github.com/repos/Codertocat/Hello-World/statuses/{sha}",
	  "languages_url": "https://api.github.com/repos/Codertocat/Hello-World/languages",
	  "stargazers_url": "https://api.github.com/repos/Codertocat/Hello-World/stargazers",
	  "contributors_url": "https://api.github.com/repos/Codertocat/Hello-World/contributors",
	  "subscribers_url": "https://api.github.com/repos/Codertocat/Hello-World/subscribers",
	  "subscription_url": "https://api.github.com/repos/Codertocat/Hello-World/subscription",
	  "commits_url": "https://api.github.com/repos/Codertocat/Hello-World/commits{/sha}",
	  "git_commits_url": "https://api.github.com/repos/Codertocat/Hello-World/git/commits{/sha}",
	  "comments_url": "https://api.github.com/repos/Codertocat/Hello-World/comments{/number}",
	  "issue_comment_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/comments{/number}",
	  "contents_url": "https://api.github.com/repos/Codertocat/Hello-World/contents/{+path}",
	  "compare_url": "https://api.github.com/repos/Codertocat/Hello-World/compare/{base}...{head}",
	  "merges_url": "https://api.github.com/repos/Codertocat/Hello-World/merges",
	  "archive_url": "https://api.github.com/repos/Codertocat/Hello-World/{archive_format}{/ref}",
	  "downloads_url": "https://api.github.com/repos/Codertocat/Hello-World/downloads",
	  "issues_url": "https://api.github.com/repos/Codertocat/Hello-World/issues{/number}",
	  "pulls_url": "https://api.github.com/repos/Codertocat/Hello-World/pulls{/number}",
	  "milestones_url": "https://api.github.com/repos/Codertocat/Hello-World/milestones{/number}",
	  "notifications_url": "https://api.github.com/repos/Codertocat/Hello-World/notifications{?since,all,participating}",
	  "labels_url": "https://api.github.com/repos/Codertocat/Hello-World/labels{/name}",
	  "releases_url": "https://api.github.com/repos/Codertocat/Hello-World/releases{/id}",
	  "deployments_url": "https://api.github.com/repos/Codertocat/Hello-World/deployments",
	  "created_at": "2018-05-30T20:18:04Z",
	  "updated_at": "2018-05-30T20:18:50Z",
	  "pushed_at": "2018-05-30T20:18:48Z",
	  "git_url": "git://github.com/Codertocat/Hello-World.git",
	  "ssh_url": "git@github.com:Codertocat/Hello-World.git",
	  "clone_url": "https://github.com/Codertocat/Hello-World.git",
	  "svn_url": "https://github.com/Codertocat/Hello-World",
	  "homepage": null,
	  "size": 0,
	  "stargazers_count": 0,
	  "watchers_count": 0,
	  "language": null,
	  "has_issues": true,
	  "has_projects": true,
	  "has_downloads": true,
	  "has_wiki": true,
	  "has_pages": true,
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
	  "login": "Codertocat",
	  "id": 21031067,
	  "node_id": "MDQ6VXNlcjIxMDMxMDY3",
	  "avatar_url": "https://avatars1.githubusercontent.com/u/21031067?v=4",
	  "gravatar_id": "",
	  "url": "https://api.github.com/users/Codertocat",
	  "html_url": "https://github.com/Codertocat",
	  "followers_url": "https://api.github.com/users/Codertocat/followers",
	  "following_url": "https://api.github.com/users/Codertocat/following{/other_user}",
	  "gists_url": "https://api.github.com/users/Codertocat/gists{/gist_id}",
	  "starred_url": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/Codertocat/subscriptions",
	  "organizations_url": "https://api.github.com/users/Codertocat/orgs",
	  "repos_url": "https://api.github.com/users/Codertocat/repos",
	  "events_url": "https://api.github.com/users/Codertocat/events{/privacy}",
	  "received_events_url": "https://api.github.com/users/Codertocat/received_events",
	  "type": "User",
	  "site_admin": false
	}
  }
`
