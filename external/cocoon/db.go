package cocoon

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewDb(config *Config) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", config.DatabaseUrl)

  if err != nil {
    return nil, err
  }

  return db, nil
}
