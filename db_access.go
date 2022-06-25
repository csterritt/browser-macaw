package main

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	//"gorm.io/driver/sqlite"      // Sqlite driver based on GGO
	"gorm.io/gorm"
	//"gorm.io/gorm"
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

type SearchTable struct {
	ID             int
	Uid            string `gorm:"type:text"`
	DatasourceName string `gorm:"type:text"`
	Title          string `gorm:"type:text"`
	Subtitle       string `gorm:"type:text"`
	Body           string `gorm:"type:text"`
	Url            string `gorm:"type:text"`
}

type SearchTableFts struct {
	Uid            string `gorm:"type:text"`
	DatasourceName string `gorm:"type:text"`
	Title          string `gorm:"type:text"`
	Subtitle       string `gorm:"type:text"`
	Body           string `gorm:"type:text"`
	Url            string `gorm:"type:text"`
}

// TableName overrides the table name used by User to `profiles`
func (SearchTable) TableName() string {
	return "search_table"
}

func (SearchTableFts) TableName() string {
	return "search_table_fts"
}

func DoQuery(query string) string {
	db, err := gorm.Open(sqlite.Open(os.Getenv("HOME")+"/.config/persistory/persistory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Flamed out trying to open the database.", err)
	}

	var entry SearchTableFts
	var output string
	res := db.Where("search_table_fts match ?", query).Take(&entry)
	if res.Error != nil {
		fmt.Printf("Non-nil res on query, error %v\n", res.Error)
	} else {
		fmt.Printf("res is %#v\n", res)
		output = fmt.Sprintf("One entry:\n\tuid '%s'\n\tTitle '%s'\n\tdatasource_name '%s'\n\tsubtitle '%s'\n\tbody '%s'\n\tURL '%s'\n",
			entry.Uid, entry.Title, entry.DatasourceName, entry.Subtitle, entry.Body, entry.Url)
		fmt.Println(output)
	}

	return output
}
