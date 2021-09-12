# GithubGo api
```
TYPES

type Commit struct {
	ShaHash string    `json:"sha"`
	Commit  GitCommit `json:"commit"`
	URL     string    `json:"url"`
}
    A Commit from github

type GitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
    The GitAuthor is the author of a git commit, with a name, email, and a date

type GitCommit struct {
	Author       GitAuthor       `json:"author"`
	Message      string          `json:"message"`
	URL          string          `json:"url"`
	CommentCount string          `json:"comment_count"`
	Verification GitVerification `json:"verification"`
}
    GitCommit is just a commit

type GitVerification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}
    The GitVerification is the verification on a commit

type Repo struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	GitURL        string `json:"git_url"`
	CloneURL      string `json:"clone_url"`
	Stargazers    int64  `json:"stargazers_count"`
	Watchers      int64  `json:"watchers_count"`
	Forks         int64  `json:"forks"`
	DefaultBranch int64  `json:"default_branch"`
}
    A Repo and all it's info

func (repo Repo) GetCommits() ([]Commit, error)
    GetCommits gets all the commits from the given repository, including the git
    verifications

type User struct {
	Username  string `json:"login"`
	ID        int64  `json:"id"`
	Avatar    string `json:"avatar_url"`
	ReposURL  string `json:"repos_url"`
	Bio       string `json:"bio"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
	Hireable  bool   `json:"hireable"`
	Blog      string `json:"blog"`
}
    User is a thing for parsing in json

func GetUser(uname string) (User, error)
    GetUser will take a username and create a new user struct using the API

func (user User) GetRepos() ([]Repo, error)
    GetRepos returns an array of repositorues from the given users repos_url

```
## Example
```
l3gacy, _ := GetUser("l3gacyb3ta")

repos, _ := l3gacy.GetRepos()

commits, _ := repos[3].GetCommits()

fmt.Println(commits[len(commits)-1].Commit)
```