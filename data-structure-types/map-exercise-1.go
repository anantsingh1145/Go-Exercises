package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":       {"shaken not_stired", "martini", "007"},
		"moneypenny_jenny": {"Intelligence", "Literature", "Computer science"},
		"no_dr":            {"cats", "ice creams", "sunsets"},
	}

	for key, value := range m {
		fmt.Println("`Key:`", key, "\n`Value:`", value)
	}
}
