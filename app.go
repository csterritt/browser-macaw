package main

import (
	"context"
	"errors"
	"fmt"

	"browser_macaw/db_access"
	"browser_macaw/debug"

	"golang.design/x/clipboard"
)

// App struct
type App struct {
	ctx context.Context
}

var clipboardInitDone bool

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
	res, err := db_access.DoQuery(query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("{\"message\": \"%s\"}", err))
	} else {
		return res, nil
	}
}

// WriteToClipboard Write the given string to the clipboard
func (a *App) WriteToClipboard(val string) {
	if !clipboardInitDone {
		// Init returns an error if the package is not available for use.
		err := clipboard.Init()
		if err != nil {
			//panic(err)
			debug.DumpStringToDebugListener(fmt.Sprintf("Unable to init clipboard: %v", err))
		} else {
			clipboardInitDone = true
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(val))
}

func (a *App) DebugOutput(msg string) {
	debug.DumpStringToDebugListener(msg)
}
