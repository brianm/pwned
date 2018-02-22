package main

import (
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/DCSO/bloom"
	"github.com/GeertJohan/go.rice"
)

var helpText = `Usage: pwned [passwords]

Checks given passwords against a bloom filter derived from
https://haveibeenpwned.com/ . The bloom filter is tuned for 
a 0.000001% false positive rate.

Passwords may be passed on the command line or on stdin.
Each argument on the command line will be treated as a
candidate, or each line from stdin will be treated as a
candidate.

Candidate passwords which match (which have been pwned)
will be printed to stdout.
`

var help = false

func main() {
	flag.BoolVar(&help, "h", false, "Print help")
	flag.Parse()

	if help {
		println(helpText)
		os.Exit(0)
	}

	box := rice.MustFindBox("data")
	bloomBytes := box.MustBytes("pwned.bloom")
	filter, err := bloom.LoadFromBytes(bloomBytes, false)
	if err != nil {
		log.Fatalf("error loading filter: %s", err)
	}

	if len(os.Args) > 1 {
		// take passwords on command line
		for _, p := range flag.Args() {
			h := sha1.New()
			io.WriteString(h, p)
			key := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
			if filter.Check([]byte(key)) {
				println(p)
			}
		}
	} else {
		// take passwords from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			p := scanner.Text()
			h := sha1.New()
			io.WriteString(h, p)
			key := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
			if filter.Check([]byte(key)) {
				println(p)
			}
		}
	}
}
