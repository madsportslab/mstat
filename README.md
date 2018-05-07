# mstat
golang library for persisting statistics, provides a simple api for counters
and logging.  also includes convenience methods to help with advanced statistics 
calculations.

originally, i used redis as the backing store, but wanted something embedded,
simple to manage, and optimized for my particular workflows so i wrote this
library on top of sqlite3.

## library usage

### Count(key, field, increment) int

adjusts counter and returns counter value as an integer.  counter values are always 0 or greater so any decrement that causes the counter to be less than 0 will result in a no-op, meaning the action will not be persisted.  counter default value should be 0.  counter max value should be 2^32 - 1.

if key doesn't exist, it will be created, if key exists then it will be updated.

keys should be unique and can be a combination of letters, numbers, or period/dot character.

```
ret := mstat.Count("game1", "1PTM", 1)

fmt.Println(ret)  // 1

```

```
ret := mstat.Count("game1", "3PTA", 2)

fmt.Println(ret)  // 2

```

```
ret1 := mstat.Count("game1", "AST", 2)
ret2 := mstat.Count("game1", "AST", -1)

fmt.Println(ret1)  // 2
fmt.Println(ret2)  // 1

```

```
ret := mstat.Count("game1", "BLK", -1)

fmt.Println(ret)  // 0

```  

### GetCounter(key) (map[string]int, error)

returns map of integers, key is string based

#### Errors

key not found

### DelCounter(key) int

deletes counter based on key, returns the number of records removed.  if the key
does not exist, a no-op is performed and 0 is returned.

### AppendLog(key, field, val) string

keys should be unique and can be a combination of letters, numbers, or  period/dot character.

appends string value to field of a key, 
`mstat.AppendLog("game1", "plays", "1PTM HOME 3")
