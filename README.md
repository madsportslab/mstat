# mstat
embedded golang library for statistics storage, leverages sqlite3 as the backing
store.  includes convenience methods to help with advanced statistics calculations.

originally, i had used redis as the backing store, but wanted something embedded,
simple to manage, and optimized for my particular workflows so i wrote a key
value library on top of sqlite3.

## library usage

### counters

counter(key, increment)

###

add(key, incr) // add("1.1.2.3", 1) add("1.1.2.3", -1)
add_set(key, field, incr)

### integer based counters

add(key, field, increment)
get(key, field)
getall(key)

1PT, 2PT, 3PT, STL, TOV, BLK, OREB, DREB, AST, FOUL  

### append log, timestamp, string

append(key, field/timestamp, increment)
play, "string"

game, team, player, 