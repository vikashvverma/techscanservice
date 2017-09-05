package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"

	"github.com/vikashvverma/techscanservice/config"
	"github.com/vikashvverma/techscanservice/factory"
	"github.com/vikashvverma/techscanservice/router"
	"github.com/vikashvverma/techscanservice/log"
)

const (
	appName = "techscan"
)

var version = "dev"

func main() {
	var c *config.Config

	var errors []error

	hasFlags := !(len(os.Args) > 2 && os.Args[1] == "-config")
	if hasFlags {
		c, errors = configFromFlag()
		if errors != nil {
			logrus.Fatalln(errors)
		}
	} else {
		path := os.Args[2]
		c, errors = configFromFile(path)
		if errors != nil {
			logrus.Fatalln(errors)
		}
	}

	f := factory.New(c, logrus.New())

	//Optionally Seed DB
	err := f.Seeder().OptionallySeedDB()
	if err != nil {
		logrus.Fatalln(err)
	}

	muxRouter := router.Router(c, f)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{c.OriginAllowed()})
	methodsOk := handlers.AllowedMethods([]string{"GET"})
	handlers.CORS(headersOk, originsOk, methodsOk)

	n := negroni.New()
	n.Use(log.New())
	n.UseHandler(handlers.CORS(headersOk, originsOk, methodsOk)(muxRouter))
	n.Run(fmt.Sprintf(":%d", c.Port()))
}

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
}

func configFromFile(path string) (*config.Config, []error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, []error{fmt.Errorf("unable to open config file '%s': %s", path, err)}
	}

	args := &config.Args{}
	err = json.Unmarshal(content, args)
	if err != nil {
		return nil, []error{fmt.Errorf("config file not valid: %s", err)}
	}

	return config.New(args)
}

func configFromFlag() (*config.Config, []error) {
	h := flag.Bool("h", false, "Help Message")

	args := &config.Args{}
	flag.StringVar(&args.Port, "port", "80", "Application Port")
	flag.StringVar(&args.DBServer, "dbServer", "", "Database Server")
	flag.StringVar(&args.DBPort, "dbPort", "5432", "Database Port")
	flag.StringVar(&args.DBName, "dbName", "", "Database Name")
	flag.StringVar(&args.DBUserName, "dbUserName", "", "Database Username")
	flag.StringVar(&args.DBPassword, "dbPassword", "", "Database Password")
	flag.StringVar(&args.DBTimeout, "dbTimeout", "30000", "Database Query Timeout in milliseconds")
	flag.StringVar(&args.SeedDataPath, "seedDataPath", "", "Seed data path, optional if data is already populated")
	flag.StringVar(&args.OriginAllowed, "originAllowed", "", "allowd origin to avoid CORS error, not required if both frontend & backend are on same domain")
	flag.Parse()

	if *h {
		flag.Usage = usage
		flag.Usage()
		os.Exit(0)
	}

	return config.New(args)
}

const (
	usagePrefix = `
Tech Scan

Usage:
    techscan -config [FILE]

Command-line flags if no config file is used:
`
)
