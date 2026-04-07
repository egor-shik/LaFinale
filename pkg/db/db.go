package db

import (
"database/sql"
"os"

_ "modernc.org/sqlite"
)

var db *sql.DB
//вроде ок, но мог накосячить
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
			return err
		}

	db = dBase

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}
	return nil
}