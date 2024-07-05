package main

import (
	"fmt"

	"github.com/jeer00/strykscraper/strykscraper"
)

func main() {
	odds, bets, err := strykscraper.FetchData()

	fmt.Println(odds, bets)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(odds)
}
