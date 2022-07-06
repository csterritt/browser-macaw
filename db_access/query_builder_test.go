package db_access

import "testing"

func TestSimpleWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build simple words query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := "foo bar baz"
	if len(args) != 1 || args[0] != expectedArgs {
		t.Errorf("Expected args '%s', got '%#v'",
			expectedArgs, args)
	}
}

func TestSimplePhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		ExactPhrase: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build simple phrase query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := "\"foo bar baz\""
	if len(args) != 1 || args[0] != expectedArgs {
		t.Errorf("Expected args '%s', got '%#v'",
			expectedArgs, args)
	}
}

func TestWordsAndPhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words:       "quux",
		ExactPhrase: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build words and phrase query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux", "\"foo bar baz\""}
	if len(args) != 2 || args[0] != expectedArgs[0] || args[1] != expectedArgs[1] {
		t.Errorf("Expected args '%s', got '%#v'",
			expectedArgs, args)
	}
}

func TestWordsAndOneUrlWordQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "quux",
		InUrl: "shazbat",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build words and url query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + UrlWhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux", "%shazbat%"}
	if len(args) != 2 || args[0] != expectedArgs[0] || args[1] != expectedArgs[1] {
		t.Errorf("Expected args '%s', got '%#v'",
			expectedArgs, args)
	}
}

func TestWordsAndManyUrlWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "quux",
		InUrl: "shazbat gronkle spurple",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build words and url query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + UrlWhereClause + And + UrlWhereClause + And + UrlWhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux", "%shazbat%", "%gronkle%", "%spurple%"}
	if len(args) != len(expectedArgs) {
		t.Errorf("Expected %d args '%s', got %d '%#v'",
			len(expectedArgs), expectedArgs, len(args), args)
	}
	for index := range expectedArgs {
		if args[index] != expectedArgs[index] {
			t.Errorf("Expected args #%d to be '%s', got '%#v'",
				index, expectedArgs[index], args[index])
		}
	}
}
