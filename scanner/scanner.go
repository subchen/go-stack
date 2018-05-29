package scanner

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

var (
	ErrEOF       = errors.New("reach at EOF")
	ErrRuneError = errors.New("utf8 rune decode error")
)

// Scanner provides a convenient interface for reading unicode data using regexp match
type Scanner struct {
	input string
	pos   int   // byte pos on input
	err   error // last error
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
	// reset
	s.err = nil

	if s.Eof() {
		s.err = ErrEOF
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
	_, ok := s.Scan(reWhitespaces)
	return ok
}

// ScanIdentifier returns an identifier
func (s *Scanner) ScanIdentifier() (string, bool) {
	return s.Scan(reIdentifier)
}

// ScanQuoteString returns an quote string
func (s *Scanner) ScanQuoteString() (string, bool) {
	if s.Eof() {
		s.err = ErrEOF
		return "", false
	}

	var re *regexp.Regexp

	peek := s.Peek()
	switch peek {
	case '"':
		re = reStringDouble
	case '\'':
		re = reStringSingle
	case '`':
		re = reStringRaw
	default:
		s.err = errors.New("no quote string found")
		return "", false
	}

	find, ok := s.Scan(re)
	if !ok {
		return "", false
	}

	unquote, err := strconv.Unquote(find)
	if err != nil {
		s.err = fmt.Errorf("failed to Unquote(), %v", err)
		return "", false
	}
	return unquote, true
}

// ScanString returns a string, the string can be quoted by '"', '\'', '`'
func (s *Scanner) ScanString() (string, bool) {
	if s.Eof() {
		s.err = ErrEOF
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

	find, ok := s.Scan(re)
	if !ok {
		return "", false
	}
	if quoted {
		unquote, err := strconv.Unquote(find)
		if err != nil {
			s.err = fmt.Errorf("failed to Unquote(), %v", err)
			return "", false
		}
		find = unquote
	}
	return find, true
}

// ScanInt64 returns an int64
func (s *Scanner) ScanInt64() (int64, bool) {
	find, ok := s.Scan(reInt)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseInt(find, 10, 64)
	if err != nil {
		s.err = fmt.Errorf("failed to ParseInt(), %v", err)
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
		s.err = fmt.Errorf("failed to ParseFloat(), %v", err)
		return 0, false
	}
	return n, true
}

// ScanChar returns next rune
func (s *Scanner) ScanChar() (rune, bool) {
	peek := s.Peek()

	if peek == EOF || peek == utf8.RuneError {
		return peek, false
	}

	s.pos += len(string(peek))
	return peek, true
}

// ScanUntil returns text before delim
func (s *Scanner) ScanUntil(delim rune) (string, bool) {
	// reset
	s.err = nil

	if s.Eof() {
		s.err = ErrEOF
		return "", false
	}

	i := s.pos
	for i < len(s.input) {
		find, size := utf8.DecodeRuneInString(s.input[i:])
		if find == utf8.RuneError {
			if size == 0 {
				s.err = ErrEOF
				return "", false
			}
			s.err = ErrRuneError
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

func (s *Scanner) ScanToEnd() string {
	// reset
	s.err = nil

	find := s.input[s.pos:]

	// move to eof
	s.pos = len(s.input)

	return find
}

// Peek returns a rune, result may be EOF, utf8.RuneError
func (s *Scanner) Peek() rune {
	// reset
	s.err = nil

	if s.Eof() {
		s.err = ErrEOF
		return EOF
	}

	find, size := utf8.DecodeRuneInString(s.input[s.pos:])
	if find == utf8.RuneError {
		if size == 0 {
			s.err = ErrEOF
			return EOF
		}
		s.err = ErrRuneError
		return utf8.RuneError
	}

	return find
}

// PeekN returns some runes, if no enough runes, return all
func (s *Scanner) PeekN(n int) []rune {
	// reset
	s.err = nil

	if s.Eof() {
		return nil
	}

	runes := make([]rune, 0, n)
	i := s.pos
	for i < len(s.input) {
		find, size := utf8.DecodeRuneInString(s.input[i:])
		if find == utf8.RuneError {
			if size == 0 {
				s.err = ErrEOF
				return runes
			}
			s.err = ErrRuneError
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

	if s.err != nil {
		msg = msg + " " + s.err.Error()
	}

	if s.Eof() {
		return fmt.Errorf("failed at EOF: %v", msg)
	}
	return fmt.Errorf("failed at %v: %v", s.pos, msg)
}
