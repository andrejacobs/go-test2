package sorted

import (
	"sort"

	"github.com/andrejacobs/go-test1/quote"
)

type QuoteList struct {
	quote.QuoteList
}

func (ql *QuoteList) Sort() {
	sort.SliceStable(ql.Quotes, func(i, j int) bool {
		return ql.Quotes[i].Line < ql.Quotes[j].Line
	})
}

func (ql *QuoteList) Add(who string, line string) {
	ql.QuoteList.Add(who, line)
	ql.Sort()
}
