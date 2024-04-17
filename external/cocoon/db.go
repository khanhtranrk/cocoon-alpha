package cocoon

import (
	"database/sql"
	"github.com/khanhtranrk/justice-alpha/external/adapter/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewDb(config *config.DB) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", config.DatabaseURL)

  if err != nil {
    return nil, err
  }

  return db, nil
}
