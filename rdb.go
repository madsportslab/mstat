package mstat

import (
	"encoding/json"
	"errors"
)

func createCounter(key string, field string, increment int) (int, error) {

	m := make(map[string]int)

	sum := m[field] + increment

	if sum > -1 {
		m[field] = sum
	} else {
		m[field] = 0
	}
	
	j, err := json.Marshal(m)

	if err != nil {
		return 0, errors.New("mstat: " + err.Error())
	} else {

		_, err := conn.Exec(
			CountCreate, key, j,
		)
	
		if err != nil {
			return 0, errors.New("mstat: " + err.Error())
		} else {
			return m[field], nil
		}
	
	}

} // createCounter


func updateCounter(key string, val string, field string, increment int) (int, error) {

	m := make(map[string]int)

	err := json.Unmarshal([]byte(val), &m)

	if err != nil {
		return 0, errors.New("mstat: " + err.Error())
	} else {

		sum := m[field] + increment

		if sum > -1 {
			m[field] = sum
		}

		j, err := json.Marshal(m)

		if err != nil {
			return 0, errors.New("mstat: " + err.Error())
		} else {

			_, err := conn.Exec(
				CountUpdate, j, key,
			)
	
			if err != nil {
				return 0, errors.New("mstat: " + err.Error())
			} else {
				return m[field], nil
			}

		}

	}

} // updateCounter
