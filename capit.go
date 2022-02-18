// capitalise from stdin to stdout
// SvM 29-DEC-2020 - 02-FEB-2021
//
// nicked from https://gobyexample.com/line-filters

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// since strings.Title only works one way (lower to upper), lower everything first
		fmt.Println(strings.Title(strings.ToLower(scanner.Text())))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
