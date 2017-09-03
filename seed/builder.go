package seed

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/vikashvverma/techscanservice/repository"
)

const (
	schema = "github"
	table  = "repository"
)

func isSeedDataPopulated(e repository.Execer) bool {
	query := fmt.Sprintf(`SELECT EXISTS (
				SELECT 1
				FROM   information_schema.tables
				WHERE  table_schema = '%s'
   				AND    table_name = '%s'
   				);`, schema, table)
	exists, err := e.Query(query, scanner)
	if err != nil {
		return false
	}

	return exists.(*struct{ id bool }).id
}

func createSchema(e repository.Execer) error {
	query := fmt.Sprintf(" CREATE SCHEMA %s;", schema)
	_, err := e.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create table: %s", err)
	}
	return nil
}
func createTable(e repository.Execer) error {
	query := fmt.Sprintf(`
	CREATE TABLE %s.%s
	(
	  repo_id 				integer
	, user_id 				integer
    , login     		VARCHAR(100)
	, avatar_url		VARCHAR(100)
	, url       		VARCHAR(100)
	, repos_url  		VARCHAR(100)
	, user_type  		VARCHAR(100)
	, html_url          VARCHAR(100)
	, language        	VARCHAR(100)
	, description       VARCHAR(300)
	, stargazers_count 	integer
	, watchers_count   	integer
	, forks_count      	integer
	);`, schema, table)

	_, err := e.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create table: %s", err)
	}
	return nil
}

func populateDB(values []string, e repository.Execer) error {
	stmt := fmt.Sprintf("INSERT INTO %s.%s (repo_id, user_id, login, avatar_url, url, repos_url, user_type, html_url, language, description, stargazers_count, watchers_count, forks_count) VALUES %s", schema, table, strings.Join(values, ","))
	_, err := e.Exec(stmt)
	return err
}

func scanner(rows *sql.Rows) (interface{}, error) {
	var result struct{ id bool }

	defer rows.Close()
	rows.Next()
	err := rows.Scan(
		&result.id,
	)
	if err != nil {
		return nil, fmt.Errorf("no row found: %s", err)
	}

	return &result, nil
}

func buildValues(ev []Repo) []string {
	var rows []string
	for _, v := range ev {
		rows = append(rows, v.string())
	}
	return rows
}
