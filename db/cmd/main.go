package main

import (
	"fmt"

	"github.com/pfremaux/golibs/db/pkg/csv"
)

func main() {
	c, err := csv.Connect("data.csv", "", "")
	if err != nil {
		panic(err)
	}
	all, err := c.ListAll("")
	if err != nil {
		panic(err)
	}
	for _, row := range all {
		fmt.Printf("%v", row)
	}
}
