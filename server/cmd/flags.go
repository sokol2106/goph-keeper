package main

import (
	"flag"
	"server/internal/config"
)

// ParseFlags
func ParseFlags(cnf *config.Config) {
	flag.StringVar(&cnf.Listen, "a", cnf.Listen, "address to run server")
	flag.StringVar(&cnf.Postgres.Host, "pgh", cnf.Postgres.Host, "file storage path")
	flag.StringVar(&cnf.Postgres.Port, "pgp", cnf.Postgres.Port, "file storage path")
	flag.StringVar(&cnf.Postgres.User, "pgu", cnf.Postgres.User, "file storage path")
	flag.StringVar(&cnf.Postgres.Password, "pgpass", cnf.Postgres.Password, "file storage path")
	flag.StringVar(&cnf.Postgres.Database, "pgdb", cnf.Postgres.Database, "file storage path")

}
