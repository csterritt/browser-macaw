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

type SearchTableFts struct {
	Uid            string `gorm:"type:text"`
	DatasourceName string `gorm:"type:text"`
	Title          string `gorm:"type:text"`
	Subtitle       string `gorm:"type:text"`
	Body           string `gorm:"type:text"`
	Url            string `gorm:"type:text"`
}

type SearchTableFtsSubset struct {
	Uid      string `gorm:"type:text"`
	Title    string `gorm:"type:text"`
	Subtitle string `gorm:"type:text"`
	Url      string `gorm:"type:text"`
}

func (SearchTableFts) TableName() string {
	return "search_table_fts"
}

func DoQuery(query string) []SearchTableFtsSubset {
	db, err := gorm.Open(sqlite.Open(os.Getenv("HOME")+"/.config/persistory/persistory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Flamed out trying to open the database.", err)
	}

	output := make([]SearchTableFtsSubset, 0)
	fullOutput := make([]SearchTableFts, 0)
	res := db.Where("search_table_fts match ?", query).
		Order("rank").Limit(50).Find(&fullOutput)
	if res.Error != nil {
		fmt.Printf("Non-nil res on query, error %v\n", res.Error)
	} else {
		output = make([]SearchTableFtsSubset, len(fullOutput))
		for index := 0; index < len(output); index++ {
			entry := fullOutput[index]
			fmt.Printf("Uid '%s'\n\tTitle '%s'\n\tSubtitle '%s'\n\tURL '%s'\n",
				entry.Uid, entry.Title, entry.Subtitle, entry.Url)

			output[index].Uid = entry.Uid
			output[index].Title = entry.Title
			output[index].Subtitle = entry.Subtitle
			output[index].Url = entry.Url
		}
	}

	return output
}
