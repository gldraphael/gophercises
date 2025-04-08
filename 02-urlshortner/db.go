package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetRulesFromDb(dbPath string) (RedirectionConfig, error) {
	db, err := getDb(dbPath)
	if err != nil {
		return RedirectionConfig{}, err
	}

	const QUERY_SQL = `
	  SELECT slug, url FROM rules;
	`
	rows, err := db.Query(QUERY_SQL)
	if err != nil {
		return RedirectionConfig{}, err
	}
	
	config := RedirectionConfig{}
	for rows.Next() {
		var rule RedirectRule
		err = rows.Scan(&rule.Slug, &rule.Url)
		if err != nil {
			return RedirectionConfig{}, err
		}
		config.Redirects = append(config.Redirects, rule)
	}
	return config, nil
}

func getDb(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	err = createDb(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDb(db *sql.DB) (error) {
	const SETUP_DB_SQL = `
		BEGIN;
	  CREATE TABLE IF NOT EXISTS rules (
      slug  TEXT NOT NULL CONSTRAINT "pk_rules" PRIMARY KEY,
      url   TEXT NOT NULL
	  );
		COMMIT;
		BEGIN;
		INSERT OR IGNORE INTO rules(slug, url) VALUES ('/dogs', 'https://bsky.app/search?q=dog+pics');
    INSERT OR IGNORE INTO rules(slug, url) VALUES ('/cats', 'https://bsky.app/search?q=cat+pics');
		COMMIT;
	`

	_, err := db.Exec(SETUP_DB_SQL)
	return err
}

