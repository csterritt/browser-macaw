package db_access

import (
	"database/sql"
	"log"
	"os"

	"browser_macaw/domain_finder"

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

type SearchTableFtsSubset struct {
	Uid        string
	Title      string
	Subtitle   string
	Url        string
	DomainName string
	BodyPart   string
}

func DoQuery(query string) []SearchTableFtsSubset {
	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.config/persistory/persistory.db")
	if err != nil {
		log.Fatal("Flamed out trying to open the database: ", err)
	}

	output := make([]SearchTableFtsSubset, 0)
	rows, err := db.Query("SELECT * FROM search_table_fts WHERE search_table_fts match ?", query)
	if err != nil {
		log.Fatal("Cannot query, got error: ", err)
	}
	defer rows.Close()

	output = make([]SearchTableFtsSubset, 0)
	seen := make(map[string]bool)
	domainsInRankOrder := make([]string, 0)
	resultIndexByDomain := make(map[string][]int)

	outputCount := 0
	for rows.Next() {
		var entry SearchTableFts
		var newEntry SearchTableFtsSubset
		err := rows.Scan(&entry.Uid, &entry.DatasourceName, &entry.Title, &entry.Subtitle, &entry.Body, &entry.Url)
		if err != nil {
			log.Fatal("Cannot scan into entry: ", err)
		}

		if entry.Uid.Valid {
			newEntry.Uid = entry.Uid.String
		}

		if entry.Title.Valid {
			newEntry.Title = entry.Title.String
		}

		if entry.Subtitle.Valid {
			newEntry.Subtitle = entry.Subtitle.String
		}

		if entry.Url.Valid {
			if _, found := seen[entry.Url.String]; found {
				continue
			}

			newEntry.Url = entry.Url.String
			seen[entry.Url.String] = true
			newEntry.DomainName = domain_finder.DomainFromUrl(entry.Url.String)

			if _, found := resultIndexByDomain[newEntry.DomainName]; !found {
				domainsInRankOrder = append(domainsInRankOrder, newEntry.DomainName)
				resultIndexByDomain[newEntry.DomainName] = make([]int, 0)
			}

			resultIndexByDomain[newEntry.DomainName] = append(resultIndexByDomain[newEntry.DomainName], outputCount)
			outputCount++

			output = append(output, newEntry)
		}
	}

	finalOutput := make([]SearchTableFtsSubset, outputCount)
	nextIndex := 0
	for _, domain := range domainsInRankOrder {
		for _, index := range resultIndexByDomain[domain] {
			finalOutput[nextIndex] = output[index]
			nextIndex++
		}
	}

	return finalOutput
}
