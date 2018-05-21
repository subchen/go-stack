package scanner

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// golang regexp doc: https://github.com/google/re2/wiki/Syntax
const (
	sWhitespaces = `[[:space:]]+`

	sIdentifier = `[_a-zA-Z][_a-zA-Z0-9]*`

	sStringDouble = `"` + sSTR + `"`
	sStringSingle = `'` + sSTR + `'`
	sStringRaw    = "`" + "(?s:[^`]*)" + "`"
	sSTR          = `(` + sESC + `|` + sOTHER + `)*?`
	sESC          = `\\([btnfr'"\\]|u[0-9a-fA-F]{4})`
	sOTHER        = `[^\r\n\\]`

	sStringPlain = `[[:graph:]]+`

	sInt   = sINT
	sFloat = sINT + `(` + `\.` + sFRAC + `)?` + sEXP + `?`
	sINT   = `(0|[1-9][0-9]*)`
	sFRAC  = `([0-9]+)`
	sEXP   = `(` + `[Ee][+\-]?` + sINT + `)`
)

var (
	reWhitespaces  = regexp.MustCompile(`^` + sWhitespaces)
	reIdentifier   = regexp.MustCompile(`^` + sIdentifier)
	reStringDouble = regexp.MustCompile(`^` + sStringDouble)
	reStringSingle = regexp.MustCompile(`^` + sStringSingle)
	reStringRaw    = regexp.MustCompile(`^` + sStringRaw)
	reStringPlain  = regexp.MustCompile(`^` + sStringPlain)
	reInt          = regexp.MustCompile(`^` + sInt)
	reFloat        = regexp.MustCompile(`^` + sFloat)
)

const (
	// EOF represents the end of input
	EOF rune = -1
)

// Scanner provides a convenient interface for reading unicode data using regexp match
type Scanner struct {
	input string
	pos   int
}

// New creates a new Scanner
func New(input string) *Scanner {
	return &Scanner{
		input: input,
		pos:   0,
	}
}

// Scan returns next matched string
func (s *Scanner) Scan(re *regexp.Regexp) (string, bool) {
	if s.Eof() {
		return "", false
	}

	find := re.FindString(s.input[s.pos:])
	if len(find) == 0 {
		return "", false
	}

	s.pos += len(find)
	return find, true
}

// SkipWhitespaces skips whitespaces, and returns true if skipped > 0
func (s *Scanner) SkipWhitespaces() bool {
	if s.Eof() {
		return false
	}

	_, ok := s.Scan(reWhitespaces)
	return ok
}

// ScanIdentifier returns an identifier
func (s *Scanner) ScanIdentifier() (string, bool) {
	return s.Scan(reIdentifier)
}

// ScanIdentifier returns a string, the string can be quoted by '"', '\'', '`'
func (s *Scanner) ScanString() (string, bool) {
	if s.Eof() {
		return "", false
	}

	var re *regexp.Regexp
	quoted := false

	peek := s.Peek()
	switch peek {
	case '"':
		re = reStringDouble
		quoted = true
	case '\'':
		re = reStringSingle
		quoted = true
	case '`':
		re = reStringRaw
		quoted = true
	default:
		re = reStringPlain
	}

	fmt.Println("re:", re)
	find, ok := s.Scan(re)
	if !ok {
		return "", false
	}
	if quoted {
		unquote, err := strconv.Unquote(find)
		if err != nil {
			return "", false
		}
		find = unquote
	}
	return find, true
}

// ScanChar returns whether next rune is matched with ch
func (s *Scanner) ScanChar(ch rune) bool {
	if s.Eof() {
		return false
	}

	find, size := utf8.DecodeRuneInString(s.input[s.pos:])
	if find != ch {
		return false
	}

	s.pos += size
	return true
}

// ScanInt64 returns an int64
func (s *Scanner) ScanInt64() (int64, bool) {
	find, ok := s.Scan(reInt)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseInt(find, 10, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

// ScanFloat64 returns an float64
func (s *Scanner) ScanFloat64() (float64, bool) {
	find, ok := s.Scan(reFloat)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseFloat(find, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

// ScanUntil returns text before delim
func (s *Scanner) ScanUntil(delim rune) (string, bool) {
	if s.Eof() {
		return "", false
	}

	for i := s.pos; i < len(s.input); i++ {
		find, size := utf8.DecodeRuneInString(s.input[i:])
		if find == utf8.RuneError {
			return "", false
		}
		if find == delim {
			findstr := s.input[s.pos:i]
			s.pos = i
			return findstr, true
		}
		i += size
	}

	return "", false
}

// Peek returns a rune, result may be EOF, utf8.RuneError
func (s *Scanner) Peek() rune {
	if s.Eof() {
		return EOF
	}

	find, size := utf8.DecodeRuneInString(s.input[s.pos:])
	if find == utf8.RuneError && size == 0 {
		return EOF
	}

	return find
}

// PeekN returns some runes, if no enough runes, return all
func (s *Scanner) PeekN(n int) []rune {
	if s.Eof() {
		return nil
	}

	runes := make([]rune, 0, n)
	for i := s.pos; i < len(s.input); i++ {
		find, size := utf8.DecodeRuneInString(s.input[i:])
		if find == utf8.RuneError {
			if size == 0 {
				return runes // EOF
			}
			// utf8 decode error
			return nil
		}
		i += size

		runes = append(runes, find)
		if len(runes) >= n {
			return runes
		}
	}

	return runes
}

// Pos returns current pos of scan
func (s *Scanner) Pos() int {
	return s.pos
}

// Eof returns whether scan to end of input
func (s *Scanner) Eof() bool {
	return s.pos >= len(s.input)
}

// Errorf reports an error with postion
func (s *Scanner) Errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	if s.Eof() {
		return fmt.Errorf("failed at EOF: %v", msg)
	}
	return fmt.Errorf("failed at %v: %v", s.pos, msg)
}
