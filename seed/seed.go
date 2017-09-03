package seed

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/vikashvverma/techscanservice/repository"
	"strings"
)

const (
	pullRequestReviewCommentEvent = "PullRequestReviewCommentEvent"
	pullRequestEvent              = "PullRequestEvent"
	forkEvent                     = "ForkEvent"
)

type Seeder interface {
	OptionallySeedDB() error
}

type DataFormat struct {
	ID   string `json:"ID"`
	Type string `json:"type"`
}

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	ReposURL  string `json:"repos_url"`
	UserType  string `json:"type"`
}

func (u *User) string() string {
	return fmt.Sprintf("%d, '%s', '%s', '%s', '%s', '%s'", u.ID, u.Login, u.AvatarURL, u.URL, u.ReposURL, u.UserType)
}

type Repo struct {
	User            User   `json:"owner"`
	HtmlURL         string `json:"html_url"`
	Language        string `json:"language"`
	Description     string `json:"description"`
	ID              int64  `json:"id"`
	StargazersCount int64  `json:"stargazers_count"`
	WatchersCount   int64  `json:"watchers_count"`
	ForksCount      int64  `json:"forks_count"`
	Forks           int64  `json:"forks"`
}

func (r Repo) string() string {
	return fmt.Sprintf("(%d, %s, '%s', '%s','%s', %d, %d, %d)", r.ID, r.User.string(), r.HtmlURL, r.Language, escapeMeta(r.Description), r.StargazersCount, r.WatchersCount, r.ForksCount)
}

type EventData struct {
	ID        string `json:"id"`
	EventType string `json:"type"`
	Payload   struct {
		PullRequest struct {
			Base struct {
				Repo Repo `json:"repo"`
			} `json:"base"`
		} `json:"pull_request"`
		Forkee Repo `json:"forkee"`
	} `json:"payload"`
}

type DBSeeder struct {
	Execer       repository.Execer
	SeedDataPath string
}

func New(e repository.Execer, seedDataPath string) Seeder {
	return &DBSeeder{Execer: e, SeedDataPath: seedDataPath}
}

func (d *DBSeeder) OptionallySeedDB() error {
	if isSeedDataPopulated(d.Execer) {
		return nil
	}

	err := createSchema(d.Execer)
	if err != nil {
		return fmt.Errorf("OptionallySeedDB: schema creation error: %s", err)
	}

	err = createTable(d.Execer)
	if err != nil {
		return fmt.Errorf("OptionallySeedDB: table creation error: %s", err)
	}

	eventData, err := d.readSeedData()
	if err != nil {
		return err
	}

	err = populateDB(eventData, d.Execer)
	if err != nil {
		return fmt.Errorf("OptionallySeedDB: could not seed db: %s", err)
	}
	return nil
}

func (d *DBSeeder) readSeedData() ([]string, error) {
	file, err := os.Open(d.SeedDataPath)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %s", err)
	}
	defer file.Close()

	var rows []Repo
	lookup := make(map[int64]int64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ok, eventData, err := filter(scanner.Text())
		if err != nil {
			return nil, err
		}
		//println(lookup[r.ID])
		if ok {
			r := repo(eventData)
			if lookup[r.ID] == 0 {
				lookup[r.ID] = 1
				rows = append(rows, r)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error occured while reading file: %s", err)
	}
	return buildValues(rows), nil
}

func filter(line string) (bool, *EventData, error) {
	var df DataFormat
	err := json.Unmarshal([]byte(line), &df)
	if err != nil {
		return false, nil, fmt.Errorf("%s is not a valid json: %s", line, err)
	}
	if df.Type == pullRequestEvent || df.Type == pullRequestReviewCommentEvent || df.Type == forkEvent {
		var ev EventData
		err = json.Unmarshal([]byte(line), &ev)
		if err != nil {
			return false, nil, err
		}
		return true, &ev, nil
	}
	return false, nil, nil
}

func repo(e *EventData) Repo {
	if e.EventType == forkEvent {
		return e.Payload.Forkee
	}
	return e.Payload.PullRequest.Base.Repo
}

func escapeMeta(value string) string {
	replace := map[string]string{"\\": "\\\\", "'": ``, "\n": "\\n"}

	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}

	return value
}
