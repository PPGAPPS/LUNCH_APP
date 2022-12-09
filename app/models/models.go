package models

import (
	"log"
	"lunch/config"
	"lunch/internal/environment"

	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// DB is a connection to your database to be used
// throughout your application.

func init() {
	bf, err := config.FS().Open("database.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = pop.LoadFrom(bf)
	if err != nil {
		log.Fatal(err)
	}

	pop.Debug = environment.Current() == environment.Development
}

func DB() *pop.Connection {
	c, err := connect(environment.Current())
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// connect returns a connection to the database based on the environment.
func connect(env string) (*pop.Connection, error) {
	c := pop.Connections[env]
	if c == nil {
		return nil, errors.Errorf("could not find connection named %s", env)
	}

	err := c.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening connection "+env)
	}

	return c, nil

}
