package main

import (
	"context"

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
func (a *App) Query(query db_access.Query) ([]db_access.ResultsByDomain, error) {
	return db_access.DoQuery(query)
}
