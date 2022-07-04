package db_access

import (
	"database/sql"
	"fmt"
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

type Query struct {
	Words        string
	ExactPhrase  string
	MustWords    string
	MustNotWords string
	OnlyDomain   string
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

func resultsFromRows(query Query, rows *sql.Rows) []ResultsByDomain {
	output := make([]SearchTableFtsSubset, 0)
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
			words := strings.Split(strings.ToLower(query.Words), " ")
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

func DoQuery(query Query) []ResultsByDomain {
	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.config/persistory/persistory.db")
	if err != nil {
		log.Fatal("Flamed out trying to open the database: ", err)
	}

	rows, err := buildQuery(query, db)
	if err != nil {
		log.Fatal("Cannot query, got error: ", err)
	}
	defer rows.Close()

	return resultsFromRows(query, rows)
}

func buildQuery(query Query, db *sql.DB) (*sql.Rows, error) {
	words := strings.Trim(query.Words, " \t\r\n")
	hasWords := len(words) > 0
	exactPhrase := strings.Trim(query.ExactPhrase, " \t\r\n")
	hasExactPhrase := len(exactPhrase) > 0
	queryText := "SELECT * FROM search_table_fts "
	args := make([]interface{}, 0)

	if hasExactPhrase {
		if strings.Index(exactPhrase, "\"") > -1 {
			exactPhrase = strings.Replace(exactPhrase, "\"", " ", 0)
		}

		exactPhrase = "\"" + exactPhrase + "\""
	}

	fmt.Printf("hasWords %v (%s), hasExactPhrase %v (%s)\n", hasWords, words, hasExactPhrase, exactPhrase)

	if hasWords && !hasExactPhrase {
		queryText += "WHERE search_table_fts match ?"
		args = append(args, words)
	}

	if !hasWords && hasExactPhrase {
		queryText += "WHERE search_table_fts match ?"
		args = append(args, exactPhrase)
	}

	if hasWords && hasExactPhrase {
		queryText += "WHERE search_table_fts match ? AND search_table_fts match ?"
		args = append(args, words)
		args = append(args, exactPhrase)
	}

	fmt.Printf("Query text is '%s', args are '%#v'\n", queryText, args)

	return db.Query(queryText, args...)
}
