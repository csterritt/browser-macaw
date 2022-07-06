package db_access

import "testing"

func TestSimpleAnyWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"foo OR bar OR baz"}
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

func TestSimpleAllWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		AllWords: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"foo AND bar AND baz"}
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

func TestSimpleAnyAndAllWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words:    "foo bar baz",
		AllWords: "quux gronkle",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"foo OR bar OR baz", "quux AND gronkle"}
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

func TestSimplePhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		ExactPhrase: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"\"foo bar baz\""}
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

func TestAnyWordsAndPhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		AllWords:    "quux gronkle",
		ExactPhrase: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux AND gronkle", "\"foo bar baz\""}
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

func TestAllWordsAndPhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words:       "quux gronkle",
		ExactPhrase: "foo bar baz",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux OR gronkle", "\"foo bar baz\""}
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

func TestOneWordOnlyUrlQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		InUrl: "shazbat",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + UrlWhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"%shazbat%"}
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

func TestManyWordOnlyUrlQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		InUrl: "shazbat urgle burgle",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + UrlWhereClause + And + UrlWhereClause + And + UrlWhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"%shazbat%", "%urgle%", "%burgle%"}
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

func TestWordsAndOneUrlWordQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "quux",
		InUrl: "shazbat",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause + And + UrlWhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"quux", "%shazbat%"}
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

func TestWordsAndManyUrlWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words: "quux",
		InUrl: "shazbat gronkle spurple",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
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

func TestOneWordForbiddenWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words:        "foo",
		MustNotWords: "shazbat",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"foo NOT shazbat"}
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

func TestOneWordManyForbiddenWordsQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		Words:        "foo",
		MustNotWords: "shazbat gronkle spurple",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"foo NOT shazbat NOT gronkle NOT spurple"}
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

func TestOneWordForbiddenWordsAndExactPhraseQuery(t *testing.T) {
	queryText, args, err := buildQuery(Query{
		ExactPhrase:  "foo bar",
		MustNotWords: "shazbat",
	})
	if err != nil {
		t.Errorf("Got error %v trying to build query", err)
	}

	expectedQueryText := QueryPrefix + WhereClause
	if queryText != expectedQueryText {
		t.Errorf("Expected queryText '%s', got '%s'",
			expectedQueryText, queryText)
	}

	expectedArgs := []string{"\"foo bar\" NOT shazbat"}
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
