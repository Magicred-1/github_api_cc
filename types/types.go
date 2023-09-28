package types

import "time"

type Repository struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	Owner       Owner  `json:"owner"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	URL         string `json:"url"`
	ForksURL    string `json:"forks_url"`
	// Ajoutez d'autres champs selon vos besoins
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at"`
	GitURL                   string    `json:"git_url"`
	SSHURL                   string    `json:"ssh_url"`
	CloneURL                 string    `json:"clone_url"`
	SVNURL                   string    `json:"svn_url"`
	Homepage                 string    `json:"homepage"`
	Size                     int       `json:"size"`
	StargazersCount          int       `json:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count"`
	Language                 string    `json:"language"`
	HasIssues                bool      `json:"has_issues"`
	HasProjects              bool      `json:"has_projects"`
	HasDownloads             bool      `json:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki"`
	HasPages                 bool      `json:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions"`
	ForksCount               int       `json:"forks_count"`
	Archived                 bool      `json:"archived"`
	Disabled                 bool      `json:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count"`
	License                  License   `json:"license"`
	AllowForking             bool      `json:"allow_forking"`
	IsTemplate               bool      `json:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
	Topics                   []string  `json:"topics"`
	Visibility               string    `json:"visibility"`
	Forks                    int       `json:"forks"`
	OpenIssues               int       `json:"open_issues"`
	Watchers                 int       `json:"watchers"`
	DefaultBranch            string    `json:"default_branch"`
}

type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SPDXID string `json:"spdx_id"`
	URL    string `json:"url"`
}
