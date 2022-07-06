package db_access

import (
	"strings"
)

const QueryPrefix = "SELECT * FROM search_table_fts WHERE "
const WhereClause = "search_table_fts match ?"
const UrlWhereClause = "url like ?"
const And = " AND "

func buildQuery(query Query) (string, []interface{}, error) {
	words := strings.Trim(query.Words, " \t\r\n")
	hasWords := len(words) > 0
	exactPhrase := strings.Trim(query.ExactPhrase, " \t\r\n")
	hasExactPhrase := len(exactPhrase) > 0
	url := strings.Trim(query.InUrl, " \t\r\n")
	hasUrl := len(url) > 0

	queryText := QueryPrefix
	args := make([]interface{}, 0)
	haveOneAlready := false

	if hasExactPhrase {
		if strings.Index(exactPhrase, "\"") > -1 {
			exactPhrase = strings.Replace(exactPhrase, "\"", " ", 0)
		}

		exactPhrase = "\"" + exactPhrase + "\""
	}

	if hasWords {
		queryText += WhereClause
		args = append(args, words)
		haveOneAlready = true
	}

	if hasExactPhrase {
		if haveOneAlready {
			queryText += And + WhereClause
		} else {
			queryText += WhereClause
		}

		args = append(args, exactPhrase)
		haveOneAlready = true
	}

	if hasUrl {
		parts := strings.Split(url, " ")
		for _, word := range parts {
			if haveOneAlready {
				queryText += And
			}

			queryText += UrlWhereClause
			args = append(args, "%"+word+"%")
			haveOneAlready = true
		}
	}

	return queryText, args, nil
}
