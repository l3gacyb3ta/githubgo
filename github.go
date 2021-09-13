package main

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
)

// User is a thing for parsing in json
type User struct {
	Username        string    `json:"login"`
	ID              int64     `json:"id"`
	Avatar          string    `json:"avatar_url"`
	ReposURL        string    `json:"repos_url"`
	Bio             string    `json:"bio"`
	Blog            string    `json:"blog"`
	Followers       int64     `json:"followers"`
	Following       int64     `json:"following"`
	Hireable        bool      `json:"hireable"`
	IsAdmin         bool      `json:"site_admin"`
	TwitterUsername string    `json:"twitter_username"`
	Email           string    `json:"email"`
	PublicRepos     int64     `json:"public_repos"`
	PublicGists     int64     `json:"public_gists"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GetRepos returns an array of repositorues from the given users repos_url
func (user User) GetRepos() ([]Repo, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(user.ReposURL)

	if err != nil {
		return make([]Repo, 0), nil
	}

	var userRepos []Repo

	json.Unmarshal(resp.Body(), &userRepos)

	return userRepos, nil
}

// A Repo and all it's info
type Repo struct {
	Name          string    `json:"name"`
	FullName      string    `json:"full_name"`
	Owner         RepoOwner `json:"owner"`
	Description   string    `json:"description"`
	GitURL        string    `json:"git_url"`
	CloneURL      string    `json:"clone_url"`
	Stargazers    int64     `json:"stargazers_count"`
	Watchers      int64     `json:"watchers_count"`
	Forks         int64     `json:"forks"`
	DefaultBranch int64     `json:"default_branch"`
}

// GetCommits gets all the commits from the given repository, including the git verifications
func (repo Repo) GetCommits() ([]Commit, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://api.github.com/repos/" + repo.FullName + "/commits")

	if err != nil {
		return make([]Commit, 0), err
	}

	var commits []Commit
	json.Unmarshal(resp.Body(), &commits)

	return commits, nil

}

// RepoOwner is a struct for the owner of a repo in the github API
type RepoOwner struct {
	Username string `json:"login"`
	ID       int64  `json:"id"`
	ReposURL string `json:"repos_url"`
}

// GetRepos returns an array of repositorues from the given users repos_url
func (owner RepoOwner) GetRepos() ([]Repo, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(owner.ReposURL)

	if err != nil {
		return make([]Repo, 0), nil
	}

	var userRepos []Repo

	json.Unmarshal(resp.Body(), &userRepos)

	return userRepos, nil
}

// A Commit from github
type Commit struct {
	ShaHash string    `json:"sha"`
	Commit  GitCommit `json:"commit"`
	URL     string    `json:"url"`
}

// GitCommit is just a commit
type GitCommit struct {
	Author       GitAuthor       `json:"author"`
	Message      string          `json:"message"`
	URL          string          `json:"url"`
	CommentCount string          `json:"comment_count"`
	Verification GitVerification `json:"verification"`
}

// The GitVerification is the verification on a commit
type GitVerification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

// The GitAuthor is the author of a git commit, with a name, email, and a date
type GitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

// GetUser will take a username and create a new user struct using the API
func GetUser(uname string) (User, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("http://api.github.com/users/" + uname)

	if err != nil {
		return User{}, err
	}

	var user User

	json.Unmarshal(resp.Body(), &user)
	return user, nil
}

func main() {

}
