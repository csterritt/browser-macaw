package db_access

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"browser_macaw/domain_finder"

	_ "github.com/mattn/go-sqlite3"
	"github.com/microcosm-cc/bluemonday"
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

type ResultsByDomain struct {
	DomainName string
	Links      []SearchTableFtsSubset
}

var policy *bluemonday.Policy
var stopWords map[string]bool

func init() {
	policy = bluemonday.UGCPolicy()
	stopWords = map[string]bool{
		"a":   true,
		"an":  true,
		"and": true,
		"not": true,
		"of":  true,
		"or":  true,
		"the": true,
	}
}

func DoQuery(query string) []ResultsByDomain {
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

		if entry.Body.Valid {
			substring := policy.Sanitize(entry.Body.String)
			findHere := strings.ToLower(substring)

			min := len(findHere) + 1
			words := strings.Split(strings.ToLower(query), " ")
			for _, word := range words {
				if _, found := stopWords[word]; !found {
					loc := strings.Index(findHere, word)
					if loc > -1 && loc < min {
						min = loc
					}
				}
			}

			if min < len(findHere)+1 {
				substring = substring[min:len(substring)]
			}

			if len(substring) > 200 {
				substring = substring[0:200]
			}
			newEntry.BodyPart = substring
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

	finalOutput := make([]ResultsByDomain, len(domainsInRankOrder))
	for domainIndex, domain := range domainsInRankOrder {
		finalOutput[domainIndex].DomainName = domain
		finalOutput[domainIndex].Links = make([]SearchTableFtsSubset, len(resultIndexByDomain[domain]))
		for index := range resultIndexByDomain[domain] {
			finalOutput[domainIndex].Links[index] = output[resultIndexByDomain[domain][index]]
		}
	}

	return finalOutput
}
