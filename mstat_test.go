package mstat

import (
	"testing"
)

func TestConnect(t *testing.T) {

	Connect("./db/mstat.db")

	Count("stat1", "1PTM", 1)
	Count("stat1", "1PTM", 1)
	Count("stat1", "1PTM", 1)
	Count("stat1", "1PTM", -1)

	Disconnect()

} // TestConnect
