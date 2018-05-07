package mstat

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	VERSION				= "0.1"
)

const (

	CountCreate = "INSERT into counters" +
		"(key, val) " +
		"VALUES($1, $2)"

	CountRemove = "DELETE from counters " +
		"WHERE key=?"

	CountGet = "SELECT id, key, val " +
		"FROM counters WHERE key=?"

	CountUpdate = "UPDATE counters " +
		"SET val=? " +
		"WHERE key=?"
		
)

type Counter struct {
	ID			string    `json:"id"`
	Key			string		`json:"key"`
	Val     string		`json:"val"`
}

type Play struct {
	ID			string    `json:"id"`
	Key			string		`json:"key"`
	Val     string		`json:"val"`
}

var conn *sql.DB

func Connect(addr string) (error) {

	_, err := os.Stat(addr)
	
	if os.IsNotExist(err) {
		return errors.New("mstat.db does not exist, create database file.")
	} else {

		db, err := sql.Open("sqlite3", addr)
	
		if err != nil {
			log.Println(err)
			return errors.New("Unable to make connection")
		}
	
		conn = db
	
		return nil
	
	}
	
} // Connect

func Disconnect() {
	conn.Close()
} // Disconnect


func Count(key string, field string, increment int) (int, error) {

	row := conn.QueryRow(
		CountGet, key,
	)

	c := Counter{}

	err := row.Scan(&c.ID, &c.Key, &c.Val)

	if err != nil {

		if err == sql.ErrNoRows {
			return createCounter(key, field, increment)
		} else {
			return 0, errors.New("mstat: " + err.Error())
		}

	} else {
		return updateCounter(key, c.Val, field, increment)
	}

} // Count

func GetCounter(key string) (map[string]int, error) {

	row := conn.QueryRow(
		CountGet, key,
	)

	c := Counter{}

	err := row.Scan(&c.ID, &c.Key, &c.Val)

	m := make(map[string]int)

	if err != nil {
		return m, errors.New("mstat: " + err.Error())
	} else {

		err := json.Unmarshal([]byte(c.Val), &m)

		if err != nil {
			return m, errors.New("mstat: " + err.Error())
		} else {
			return m, nil
		}

	}

} // GetCounter

func DelCounter(key string) int {

	row := conn.QueryRow(
		CountGet, key,
	)

	c := Counter{}

	err := row.Scan(&c.ID, &c.Key, &c.Val)

	if err != nil {
		return 0
	} else {

		_, err := conn.Exec(
			CountRemove, key,
		)

		if err != nil {
			return 0
		} else {
			return 1
		}

	}

} // DelCounter

func AppendLog(key string, val string) (string, error) {
	return "abc", nil
} // AppendLog
