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

func TestConnectDeleteCounters(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	} else {

		deleted := DelCounter("stat1")
		
		if deleted > 1 {
			t.Fatal("keys deleted: ", deleted)
		}

		deleted2 := DelCounter("stat3")
		
		if deleted2 > 1 {
			t.Fatal("keys deleted: ", deleted2)
		}

	}

} // TestConnectDeleteCounters

func TestConnectDeleteNonExistentCounter(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	} else {

		deleted := DelCounter("statx")
		
		if deleted != 0 {
			t.Fatal("keys deleted: ", deleted)
		}

	}

} // TestConnectDeleteNonExistentCounter

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

	m, err := GetCounter("stat2")

	if err != nil {
		t.Fatal(err)
	} else {
		
		if m["1PTM"] != 0 {
			t.Fatal("Count should be 0, but received: ", m["1PTM"])
		}

	}

	Disconnect()

} // TestNegativeCount

func TestConcurrentCount(t *testing.T) {

	err := Connect("./db/mstat.db")

	if err != nil {
		t.Fatal(err)
	}

	done := make(chan bool)

	for i := 0; i < 10; i++ {

		go func() {
			r, err := Count("stat3", "1PTM", 1)
			
			if err != nil {
				t.Log(err)
			} else {
				t.Log(r)
			}
			
			done <- true

		}()
	
	}

	for j := 0; j < 10; j++ {
		<-done
	}

	m, err := GetCounter("stat3")

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(m["1PTM"])
	}

	Disconnect()

} // TestConcurrentCount

