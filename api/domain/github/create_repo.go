package github

/*
	{
	"name": "golang-tutorial",
	"description": "Golang repo",
	"homepage": "https://github.com",
	"private": false,
	"has_issues": true,
	"has_projects": true,
	"has_wiki": true
	}
*/
// CreateRepoRequest
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

/*
{
    "id": 267003909,
    "node_id": "MDEwOlJlcG9zaXRvcnkyNjcwMDM5MDk=",
    "name": "golang-tutorial",
    "full_name": "stevedesilva/golang-tutorial",
    "private": false,
    "owner": {
        "login": "stevedesilva",
        "id": 7001268,
        "node_id": "MDQ6VXNlcjcwMDEyNjg=",
        "avatar_url": "https://avatars3.githubusercontent.com/u/7001268?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/stevedesilva",
        "html_url": "https://github.com/stevedesilva",
        "followers_url": "https://api.github.com/users/stevedesilva/followers",
        "following_url": "https://api.github.com/users/stevedesilva/following{/other_user}",
        "gists_url": "https://api.github.com/users/stevedesilva/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/stevedesilva/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/stevedesilva/subscriptions",
        "organizations_url": "https://api.github.com/users/stevedesilva/orgs",
        "repos_url": "https://api.github.com/users/stevedesilva/repos",
        "events_url": "https://api.github.com/users/stevedesilva/events{/privacy}",
        "received_events_url": "https://api.github.com/users/stevedesilva/received_events",
        "type": "User",
        "site_admin": false
    },
    "html_url": "https://github.com/stevedesilva/golang-tutorial",
    "description": "Golang repo",
    "fork": false,
    "url": "https://api.github.com/repos/stevedesilva/golang-tutorial",
    "forks_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/forks",
    "keys_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/teams",
    "hooks_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/hooks",
    "issue_events_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/issues/events{/number}",
    "events_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/events",
    "assignees_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/assignees{/user}",
    "branches_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/branches{/branch}",
    "tags_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/tags",
    "blobs_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/languages",
    "stargazers_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/stargazers",
    "contributors_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/contributors",
    "subscribers_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/subscribers",
    "subscription_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/subscription",
    "commits_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/contents/{+path}",
    "compare_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/merges",
    "archive_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/downloads",
    "issues_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/issues{/number}",
    "pulls_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/labels{/name}",
    "releases_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/releases{/id}",
    "deployments_url": "https://api.github.com/repos/stevedesilva/golang-tutorial/deployments",
    "created_at": "2020-05-26T09:44:41Z",
    "updated_at": "2020-05-26T09:44:41Z",
    "pushed_at": "2020-05-26T09:44:41Z",
    "git_url": "git://github.com/stevedesilva/golang-tutorial.git",
    "ssh_url": "git@github.com:stevedesilva/golang-tutorial.git",
    "clone_url": "https://github.com/stevedesilva/golang-tutorial.git",
    "svn_url": "https://github.com/stevedesilva/golang-tutorial",
    "homepage": "https://github.com",
    "size": 0,
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
    "disabled": false,
    "open_issues_count": 0,
    "license": null,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "permissions": {
        "admin": true,
        "push": true,
        "pull": true
    },
    "allow_squash_merge": true,
    "allow_merge_commit": true,
    "allow_rebase_merge": true,
    "delete_branch_on_merge": false,
    "network_count": 0,
    "subscribers_count": 1
}
*/

// CreateRepoResponse struct
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

// RepoOwner struct
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

// RepoPermissions struct
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	IsPush  bool `json:"push"`
	IsPull  bool `json:"pull"`
}
