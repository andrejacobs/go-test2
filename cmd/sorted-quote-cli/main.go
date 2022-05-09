package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/andrejacobs/go-test2/sorted"
)

func main() {
	countFlag := flag.Int("count", 1, "The number of quotes to print on STDOUT")
	flag.Parse()

	if err := run(os.Stdout, *countFlag); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out io.Writer, count int) error {
	ql := &sorted.QuoteList{}
	ql.Generate(count)
	ql.Sort()

	for _, quote := range ql.Quotes {
		_, err := fmt.Fprintf(out, "%q - %s\n", quote.Line, quote.Who)
		if err != nil {
			return err
		}
	}

	return nil
}
