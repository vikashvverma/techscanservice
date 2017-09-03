package github

import (
	"database/sql"
	"fmt"

	"github.com/vikashvverma/techscanservice/repository"
	"sort"
)

const (
	schema      = "github"
	table       = "repository"
	repoPerPage = 9
)

func findUser(e repository.Execer, repoID int64) (User, error) {
	query := fmt.Sprintf(`SELECT user_id, login, avatar_url, url, repos_url FROM %s.%s WHERE repo_id = %d`, schema, table, repoID)
	fmt.Println(query)
	res, err := e.Query(query, userScanner)
	if err != nil {
		return User{}, fmt.Errorf("findUser: %s", err)
	}
	return res.(User), nil
}

func userScanner(rows *sql.Rows) (interface{}, error) {
	defer rows.Close()
	rows.Next()
	var user User
	err := rows.Scan(
		&user.ID,
		&user.Login,
		&user.AvatarURL,
		&user.URL,
		&user.ReposURL,
	)
	if err != nil {
		return nil, fmt.Errorf("repoScanner: scan: %s", err)
	}

	return user, nil
}

func searchLang(e repository.Execer, lang string, page int64) ([]Repository, error) {
	query := fmt.Sprintf(`SELECT repo_id, login,  stargazers_count, language, html_url, description FROM %s.%s WHERE LOWER(language) LIKE '%s' ORDER BY stargazers_count DESC OFFSET %d LIMIT %d`, schema, table, lang+"%", page*repoPerPage, repoPerPage)
	res, err := e.Query(query, repoScanner)
	if err != nil {
		return nil, fmt.Errorf("searchLang: %s", err)
	}
	return res.([]Repository), nil
}

func repoScanner(rows *sql.Rows) (interface{}, error) {
	results := []Repository{}

	defer rows.Close()
	for rows.Next() {
		var result Repository
		err := rows.Scan(
			&result.ID,
			&result.Login,
			&result.StarCount,
			&result.Language,
			&result.RepoURL,
			&result.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("repoScanner: scan: %s", err)
		}
		results = append(results, result)
	}

	return results, nil
}

func findAll(e repository.Execer) ([]Language, error) {
	query := fmt.Sprintf(`SELECT repo_id, stargazers_count, forks_count, watchers_count, language FROM %s.%s`, schema, table)
	res, err := e.Query(query, languageScanner)
	if err != nil {
		return nil, fmt.Errorf("findAll: %s", err)
	}
	langs := res.([]Language)
	sort.Sort(languages(langs))
	return langs, nil
}

func languageScanner(rows *sql.Rows) (interface{}, error) {
	var results []Language

	defer rows.Close()
	for rows.Next() {
		var result Language
		err := rows.Scan(
			&result.ID,
			&result.StartCount,
			&result.ForksCount,
			&result.WatchersCount,
			&result.Language,
		)
		if err != nil {
			return nil, fmt.Errorf("languageScanner: scan: %s", err)
		}
		results = append(results, result)
	}

	return results, nil
}
