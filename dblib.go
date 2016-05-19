package dblib

import (
	"database/sql"

	sc "github.com/jsmith312/soundcloud-api"
	//go-sqlite3 sqlite library
	_ "github.com/mattn/go-sqlite3"
)

//InitDB initializes the DB
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

//CreateTable creates the group table
func CreateTable(db *sql.DB) {
	// create table if not exists
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
		Id INTEGER NOT NULL PRIMARY KEY,
		Name TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}

//StoreItem stores items into the DB
func StoreItem(db *sql.DB, groups []sc.Group) {
	sqlAdditem := `
	INSERT OR REPLACE INTO items(
		Id,
		Name,
		InsertedDatetime
	) values(?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(sqlAdditem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, group := range groups {
		_, err2 := stmt.Exec(group.ID, group.Name)
		if err2 != nil {
			panic(err2)
		}
	}
}

//ReadItem reads the item table
func ReadItem(db *sql.DB) []sc.Group {
	sqlReadall := `
	SELECT Id, Name FROM items
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sqlReadall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []sc.Group
	for rows.Next() {
		group := sc.Group{}
		err2 := rows.Scan(&group.ID, &group.Name)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, group)
	}
	return result
}

//DeleteTable deletes everything inside of the table.
func DeleteTable(db *sql.DB) {
	// create table if not exists
	sqlTable := `DELETE FROM items;`
	_, err := db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}

//DropTable deletes everything inside of the table.
func DropTable(db *sql.DB) {
	// create table if not exists
	sqlTable := `DROP TABLE IF EXISTS items;`
	_, err := db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}
