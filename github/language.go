package github

import (
	"github.com/vikashvverma/techscanservice/repository"
	"fmt"
)

type Fetcher interface {
	Fetch() ([]Repos, error)
	Language(language string, page int64) (repositories []Repository, err error)
	User(repoID int64) (User, error)
}

type topLangugae struct {
	Execer repository.Execer
}

func New(e repository.Execer) Fetcher {
	return &topLangugae{Execer: e}
}

func (t *topLangugae) Fetch() ([]Repos, error) {
	langs,err:=findAll(t.Execer)
	if err !=nil{
		return nil, fmt.Errorf("Fetch: %s", err)
	}
	return choose(langs), nil
}

func (t *topLangugae) Language(language string, page int64) (repositories []Repository, err error) {
	return searchLang(t.Execer, language, page)
}

func (t *topLangugae) User(repoID int64) (User, error) {
	return findUser(t.Execer, repoID)
}
