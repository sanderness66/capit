// capitalise from stdin to stdout
// SvM 29-DEC-2020 - 27-JUN-2022
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
	"log"
	"os"
	"path"
)

func b2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func main() {

	moi := path.Base(os.Args[0])
	lg := log.New(os.Stderr, moi+": ", 0)

	var u, l, c bool
	var lng string

	flag.BoolVar(&u, "u", false, "convert to uppercase")
	flag.BoolVar(&l, "l", false, "convert to lowercase")
	flag.BoolVar(&c, "c", false, "capitalise (default)")
	flag.StringVar(&lng, "L", "en", "select language")
	flag.Parse()

	lngtag := language.Make(lng)

	if x := b2i(u) + b2i(l) + b2i(c); x == 0 {
		c = true
	} else if x != 1 {
		lg.Printf("use one and only one of -u, -l or -c\n")
		os.Exit(1)
	}

	var cs cases.Caser
	if u {
		cs = cases.Upper(lngtag)
	} else if l {
		cs = cases.Lower(lngtag)
	} else if c {
		cs = cases.Title(lngtag)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(cs.String(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
