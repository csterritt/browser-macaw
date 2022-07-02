package main

import (
	"context"
	"strings"

	"browser_macaw/db_access"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Query queries the database on the user's behalf
func (a *App) Query(query db_access.Query) []db_access.ResultsByDomain {
	var output []db_access.ResultsByDomain
	if len(query.Words) != 0 && len(strings.Trim(query.Words, " \t\r\n")) != 0 {
		output = db_access.DoWordsQuery(query)
	} else if len(query.ExactPhrase) != 0 && len(strings.Trim(query.ExactPhrase, " \t\r\n")) != 0 {
		output = db_access.DoExactPhraseQuery(query)
	}

	return output
}
