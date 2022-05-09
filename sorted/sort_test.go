package sorted_test

import (
	"testing"

	"github.com/andrejacobs/go-test2/sorted"
)

func TestSort(t *testing.T) {
	ql := &sorted.QuoteList{}

	ql.Add("p1", "Bravo")
	ql.Add("p2", "Zebra")
	ql.Add("p3", "Alpha")

	if ql.Quotes[0].Line != "Alpha" &&
		ql.Quotes[1].Line != "Bravo" &&
		ql.Quotes[2].Line != "Zebra" {
		t.Fatal("Expected the quotes to be sorted")
	}
}
