package github

import "sort"

type Language struct {
	ID            int64  `json:"id"`
	StartCount    int64  `json:"startCount"`
	ForksCount    int64  `json:"forksCount"`
	WatchersCount string `json:"watchersCount"`
	Language      string `json:"language"`
}

type languages []Language

func (l languages) Len() int {
	return len(l)
}

func (l languages) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l languages) Less(i, j int) bool {
	return l[i].StartCount > l[j].StartCount
}

type Repository struct {
	ID          int64  `json:"id"`
	Login       string `json:"login"`
	StarCount   int64  `json:"starCount"`
	Language    string `json:"language"`
	RepoURL     string `json:"repoUrl"`
	Description string `json:"description"`
}

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	URL       string `json:"url"`
	ReposURL  string `json:"reposUrl"`
	UserType  string `json:"type"`
}

type Repos struct {
	RepoCount int64  `json:"repoCount"`
	Language  string `json:"language"`
}

type repos []Repos

func (l repos) Len() int {
	return len(l)
}

func (l repos) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l repos) Less(i, j int) bool {
	return l[i].RepoCount > l[j].RepoCount
}

func choose(langs []Language) []Repos {
	repoCount := make(map[string]int64, 0)
	for _, v := range langs {
		if _, ok := repoCount[v.Language]; !ok {
			repoCount[v.Language] = 0
		}
		repoCount[v.Language] += 1
	}

	var repositories []Repos
	for k, v := range repoCount {
		repositories = append(repositories, Repos{Language: k, RepoCount: v})
	}

	sort.Sort(repos(repositories))
	return repositories
}
