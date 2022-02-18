// capitalise from stdin to stdout
// SvM 29-DEC-2020 - 15-JUL-2021
//
// nicked from https://gobyexample.com/line-filters
// and from https://pkg.go.dev/golang.org/x/text@v0.3.6/cases

package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

func main() {
	lng := "en"
	flag.StringVar(&lng, "L", "en", "select language")
	flag.Parse()
	lngtag := language.Make(lng)
	c := cases.Title(lngtag)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(c.String(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
