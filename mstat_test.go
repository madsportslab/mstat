package mstat

import (
	"testing"
)

func TestNoFileConnect(t *testing.T) {

	err := Connect("./db/m.db")

	if err == nil {
		t.Fatal("Database file should not exist")
		Disconnect()
	}

} // TestNoFileConnect

func TestConnectDeleteCounter(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	} else {

		deleted := DelCounter("stat1")
		
		if deleted > 1 || deleted < 0 {
			t.Fatal("keys deleted: ", deleted)
		}

	}

} // TestNoFileConnect

func TestCount(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	} else {

		Count("stat1", "1PTM", 1)
		Count("stat1", "1PTM", 1)
		Count("stat1", "1PTM", 1)
		Count("stat1", "1PTM", -1)
	
		m, err := GetCounter("stat1")

		if err != nil {
			t.Fatal(err)
		} else {

			if m["1PTM"] != 2 {
				t.Fatal("count should be 2, but received: ", m["1PTM"])
			}

		}

		Disconnect()
	
	}


} // TestCount

func TestNegativeCount(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	}

	Count("stat2", "1PTM", -1)

	Disconnect()

} // TestNegativeCount

