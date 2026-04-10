package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	install := false

	if err != nil {
		install = true
	}

	dBase, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("init db error")
	}

	DB = dBase

	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			dBase.Close()
			return fmt.Errorf("init db error")
		}
	}
	return nil
}
