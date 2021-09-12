# GithubGo api
```
l3gacy, _ := GetUser("l3gacyb3ta")

repos, _ := l3gacy.GetRepos()

commits, _ := repos[3].GetCommits()

fmt.Println(commits[len(commits)-1].Commit)
```