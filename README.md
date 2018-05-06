# mstat
golang library for persisting statistics, provides a simple api for counters
and logging.  also includes convenience methods to help with advanced statistics 
calculations.

originally, i used redis as the backing store, but wanted something embedded,
simple to manage, and optimized for my particular workflows so i wrote this
library on top of sqlite3.

## library usage

### Count(key, field, increment)

`mstat.Count("game1", "1PTM", 1)`
`mstat.Count("game1", "3PTA", 2)`

### AppendLog(key, field, val)

`mstat.AppendLog("game1", "plays", "1PTM HOME 3")
