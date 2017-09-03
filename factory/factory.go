package factory

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq" //PostgreSQL driver

	"github.com/vikashvverma/techscanservice/config"
	"github.com/vikashvverma/techscanservice/github"
	"github.com/vikashvverma/techscanservice/repository"
	"github.com/vikashvverma/techscanservice/seed"
)

type Factory struct {
	config *config.Config
	conn   *sql.DB
	logger *logrus.Logger
}

func New(c *config.Config, l *logrus.Logger) *Factory {
	dmDB, err := newDatabase("postgres", c.ConnectionString())
	if err != nil {
		log.Fatal("Could not establish connection:", err)
	}
	return &Factory{
		config: c,
		logger: l,
		conn:   dmDB,
	}
}

func (f *Factory) Logger() *logrus.Logger {
	return f.logger
}

func (f *Factory) Seeder() seed.Seeder {
	return seed.New(repository.New(f.conn), f.config.SeedDataPath())
}

func (f *Factory) Fetcher() github.Fetcher {
	return github.New(repository.New(f.conn))
}

func newDatabase(driverName, connectionString string) (*sql.DB, error) {
	sqlDB, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to open %s: %s", driverName, err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping %s sqlDB: %s", driverName, err)
	}

	return sqlDB, nil
}
