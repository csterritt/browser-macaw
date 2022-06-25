package main

import (
	"context"
	"strings"
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
func (a *App) Query(query string) string {
	output := ""
	if len(query) != 0 && len(strings.Trim(query, " \t\r\n")) != 0 {
		output = DoQuery(query)
	} else {
		output = "Empty string given to query."
	}

	return output
}
