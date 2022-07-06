package db_access

import (
	"strings"
)

const QueryPrefix = "SELECT * FROM search_table_fts WHERE "
const WhereClause = "search_table_fts match ?"
const UrlWhereClause = "url like ?"
const And = " AND "
const Not = " (NOT "

func cleanArgAndFlag(param string) (string, bool) {
	res := strings.Trim(param, " \t\r\n")
	return res, res != ""
}

func buildQuery(query Query) (string, []interface{}, error) {
	words, hasWords := cleanArgAndFlag(query.Words)
	exactPhrase, hasExactPhrase := cleanArgAndFlag(query.ExactPhrase)
	url, hasUrl := cleanArgAndFlag(query.InUrl)
	//mustAppear, hasMustAppear := cleanArgAndFlag(query.MustWords)
	mustNotAppear, hasMustNotAppear := cleanArgAndFlag(query.MustNotWords)

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

	if hasMustNotAppear {
		last := len(args) - 1
		args[last] = query.Words + " NOT " + mustNotAppear
		haveOneAlready = true
	}

	//fmt.Printf("Built query is '%s', args are '%#v'\n", queryText, args)

	return queryText, args, nil
}
