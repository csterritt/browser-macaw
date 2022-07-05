package db_access

import (
	"strings"
)

const QueryPrefix = "SELECT * FROM search_table_fts "
const WhereClause = "WHERE search_table_fts match ? "
const PartialWhereClause = "AND search_table_fts match ?"

func buildQuery(query Query) (string, []interface{}, error) {
	words := strings.Trim(query.Words, " \t\r\n")
	hasWords := len(words) > 0
	exactPhrase := strings.Trim(query.ExactPhrase, " \t\r\n")
	hasExactPhrase := len(exactPhrase) > 0
	queryText := QueryPrefix
	args := make([]interface{}, 0)

	if hasExactPhrase {
		if strings.Index(exactPhrase, "\"") > -1 {
			exactPhrase = strings.Replace(exactPhrase, "\"", " ", 0)
		}

		exactPhrase = "\"" + exactPhrase + "\""
	}

	if hasWords && !hasExactPhrase {
		queryText += WhereClause
		args = append(args, words)
	}

	if !hasWords && hasExactPhrase {
		queryText += WhereClause
		args = append(args, exactPhrase)
	}

	if hasWords && hasExactPhrase {
		queryText += WhereClause + PartialWhereClause
		args = append(args, words)
		args = append(args, exactPhrase)
	}

	return queryText, args, nil
}
