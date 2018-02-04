package ss

import (
	"fmt"
	"strings"
)

type quoteChar struct {
	begin byte
	end   byte
}

// ",',[]
func newQuoteChars(chars string) []*quoteChar {
	cs := strings.Split(chars, ",")

	quoteCharList := make([]*quoteChar, 0, len(cs))
	for _, c := range cs {
		c = strings.TrimSpace(c)
		if len(c) == 1 {
			quoteCharList = append(quoteCharList, &quoteChar{
				begin: c[0],
				end:   c[0],
			})
		} else if len(c) > 1 {
			quoteCharList = append(quoteCharList, &quoteChar{
				begin: c[0],
				end:   c[1],
			})
		}
	}
	return quoteCharList
}

// SplitWithQuotes splits strings by given separator except treating quoted part as a single token
// quoteChars samples: `',"`, `',",[],(),[]`
func SplitWithQuotes(input, sep string, quoteChars string, keepQuotes bool) []string {
	if sep == "" {
		sep = " "
	}

	quoteCharList := newQuoteChars(quoteChars)
	fmt.Println(quoteCharList)

	var quoteEnd byte
	var buf []string
	var ret []string

	items := strings.Split(input, sep)
	for _, s := range items {
		if buf != nil {
			buf = append(buf, s)

			// isEnd ?
			ss := strings.TrimSpace(s)
			if ss[len(ss)-1] == quoteEnd {
				sss := strings.TrimSpace(strings.Join(buf, sep))
				if !keepQuotes {
					sss = sss[1 : len(sss)-1]
				}
				ret = append(ret, sss)
				buf = nil
				continue
			}
		}

		ss := strings.TrimSpace(s)
		if ss == "" {
			continue
		}

		quote := false
		for _, q := range quoteCharList {
			if q.begin == ss[0] {
				if q.end == ss[len(ss)-1] {
					sss := ss
					if !keepQuotes {
						sss = ss[1 : len(ss)-1]
					}
					ret = append(ret, sss)
				} else {
					buf = append(buf, s)
					quoteEnd = q.end
				}

				quote = true
				break
			}
		}

		if !quote {
			ret = append(ret, ss)
		}
	}

	// no close quote char
	if buf != nil {
		for _, s := range buf {
			ss := strings.TrimSpace(s)
			if ss != "" {
				ret = append(ret, ss)
			}
		}
	}

	return ret
}
