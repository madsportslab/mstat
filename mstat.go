package mstat

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	VERSION				= "0.1"
)

const (

	CountCreate = "INSERT into stats" +
		"(key, val) " +
		"VALUES($1, $2)"

	CountGet = "SELECT id, key, val " +
		"FROM stats WHERE key=?"

	CountUpdate = "UPDATE stats " +
		"SET val=? " +
		"WHERE key=?"
		
)

type Data struct {
	ID			string    `json:"id"`
	Key			string		`json:"key"`
	Val     string		`json:"val"`
}

var conn *sql.DB

func add(field string, increment int) string {

	m := make(map[string]int)

	m[field] = m[field] + increment

	j, err := json.Marshal(m)

	if err != nil {
		
		log.Println(err)
		return ""

	} else {
		return string(j)
	}

} // add

func update(j string, field string, increment int) string {

	m := make(map[string] int)
	
	err := json.Unmarshal([]byte(j), &m)

	if err != nil {
		
		log.Println(err)
		return ""

	} else {

		m[field] = m[field] + increment

		j_updated, err := json.Marshal(m)
	
		if err != nil {
			log.Println(err)
		}

		return string(j_updated)
	
	}

} // update

func Connect(addr string) {

	db, err := sql.Open("sqlite3", addr)
	
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	conn = db
	
} // Connect

func Disconnect() {
	conn.Close()
} // Disconnect

func Count(key string, field string, increment int) {

	row := conn.QueryRow(
		CountGet, key,
	)

	d := Data{}

	err := row.Scan(&d.ID, &d.Key, &d.Val)

	if err != nil {
		log.Println(err)

		if err == sql.ErrNoRows {

			_, err := conn.Exec(
				CountCreate, key, add(field, increment),
			)

			if err != nil {
				log.Println(err)
			}

		}

	} else {

		_, err := conn.Exec(
			CountUpdate, update(d.Val, field, increment), key,
		)

		if err != nil {
			log.Println(err)
		}

	}

} // Count

func AppendLog(key string, val string) {

} // AppendLog
