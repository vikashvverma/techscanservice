package config

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Config struct {
	port               int
	dbConnectionString string
	seedDataPath       string
	originAllowed      string
}

type Args struct {
	Port          string `json:"port"`
	DBName        string `json:"dbName"`
	DBServer      string `json:"dbServer"`
	DBPort        string `json:"dbPort"`
	DBUserName    string `json:"dbUserName"`
	DBPassword    string `json:"dbPassword"`
	DBTimeout     string `json:"dbTimeout"`
	SeedDataPath  string `json:"seedDataPath"`
	OriginAllowed string `json:"originAllowed"`
}

func New(args *Args) (*Config, []error) {
	c := &Config{}

	var err error
	var errors []error

	c.port = 80
	if args.Port != "" {
		c.port, err = strconv.Atoi(args.Port)
		if err != nil {
			errors = append(errors, fmt.Errorf("invalid value %q for port: %s", args.Port, err))
		}
	}

	c.seedDataPath = args.SeedDataPath

	err = validate(args)
	if err != nil {
		log.Fatalf("Config initialization failed: %s", err)
	}

	c.dbConnectionString, err = connectionString(args.DBName, args.DBServer, args.DBUserName, args.DBPassword, args.DBPort, args.DBTimeout)
	if err != nil {
		errors = append(errors, err)
	}

	if errors != nil && len(errors) > 0 {
		return nil, errors
	}

	c.originAllowed = args.OriginAllowed
	return c, nil
}

func validate(args *Args) error {
	missing := []string{}
	if args.DBPort == "" {
		missing = append(missing, "dbPort")
	}
	if args.DBServer == "" {
		missing = append(missing, "dbServer")
	}
	if args.DBName == "" {
		missing = append(missing, "dbName")
	}
	if args.DBUserName == "" {
		missing = append(missing, "dbUserName")
	}
	if args.DBPassword == "" {
		missing = append(missing, "dbPassword")
	}

	if len(missing) > 0 {
		return fmt.Errorf("%s not found", strings.Join(missing, ", "))
	}

	return nil
}
func connectionString(dbName, dbServer, dbUserName, dbPassword, dbPort, dbTimeout string) (string, error) {
	connection := fmt.Sprintf(
		"postgres://%s:%s/%s?sslmode=disable&statement_timeout=%s",
		dbServer,
		dbPort,
		dbName,
		dbTimeout,
	)

	u, err := url.Parse(connection)
	if err != nil {
		return "", fmt.Errorf("failed to build connection string: %s", err)
	}
	u.User = url.UserPassword(dbUserName, dbPassword)
	return u.String(), nil
}

func (c *Config) ConnectionString() string {
	return c.dbConnectionString
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) SeedDataPath() string {
	return c.seedDataPath
}

func (c *Config) OriginAllowed() string {
	return c.originAllowed
}
