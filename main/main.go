package main

import (
	"fmt"
)

const (
	MARQUISE = 1
	EYRIE = 2
	WOODLAND = 3
	VAGABOND = 4
)

var roles = map[int]string{
	MARQUISE: "Marquise de Cat",
	EYRIE:    "Eyrie Dynasty",
	WOODLAND: "Woodland Alliance",
	VAGABOND: "Vagabond",
}

var people = map[string]string{
	"Brad":  "",
	"Kevin": roles[WOODLAND],
	"Jacob": roles[EYRIE],
	"Jacob's Bro":  roles[MARQUISE],
}

func main() {
	switch len(people) {
	case 2:
		delete(roles, WOODLAND)
		delete(roles, VAGABOND)
	case 3:
		delete(roles, VAGABOND)
	}

	for name, prevRole := range people {
		for key, role := range roles{
			if role != prevRole {
				fmt.Printf("%s: %s\n", name, role)
				delete(roles, key)
				break
			}
		}
	}
}