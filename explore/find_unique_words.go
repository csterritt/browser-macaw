package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

/*
CREATE TABLE search_table (
                       id integer primary key,
                       Uid text,
                       datasource_name text,
                       Title text,
                       Subtitle text,
                       Body text,
                       Url text
                      );

	CREATE VIRTUAL TABLE search_table_fts using fts5(
		uid unindexed,
		datasource_name unindexed,
		title,
		subtitle,
		body,
		url,
		content='search_table',
		content_rowid='id',
		tokenize='unicode61',
		prefix='2 3 4 5'
	 )
*/

type SearchTableFts struct {
	Uid            sql.NullString
	DatasourceName sql.NullString
	Title          sql.NullString
	Subtitle       sql.NullString
	Body           sql.NullString
	Url            sql.NullString
}

func findUniqueWords() {
	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.config/persistory/persistory.db")
	if err != nil {
		log.Fatal("Flamed out trying to open the database: ", err)
	}

	rows, err := db.Query("SELECT title, subtitle, body FROM search_table_fts")
	if err != nil {
		log.Fatal("Cannot run query, got error: ", err)
	}
	defer rows.Close()

	output := make(map[string]int)
	count := 0
	for rows.Next() {
		var entry SearchTableFts
		err := rows.Scan(&entry.Title, &entry.Subtitle, &entry.Body)
		if err != nil {
			log.Fatal("Cannot scan into entry: ", err)
		}

		if entry.Title.Valid {
			for _, part := range strings.Split(entry.Title.String, " ") {
				part = strings.ToLower(part)
				output[part] += 1
			}
		}

		if entry.Subtitle.Valid {
			for _, part := range strings.Split(entry.Subtitle.String, " ") {
				part = strings.ToLower(part)
				output[part] += 1
			}
		}

		if entry.Body.Valid {
			for _, part := range strings.Split(entry.Body.String, " ") {
				part = strings.ToLower(part)
				output[part] += 1
			}
		}

		count++
	}
	fmt.Println("Found", count, "rows.")

	for word, count := range output {
		if count == 1 {
			fmt.Println(word)
		}
	}
}

func main() {
	findUniqueWords()
}
