package unpack

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Unpack(str string) (string, error) {
	if !utf8.ValidString(str) {
		return "", NewErrNotValidUtf8()
	}
	var (
		sb         strings.Builder
		lastCh     = utf8.RuneError
		pos        int
		r          rune
		strOffset  int
		isEscaping bool
	)
	for len(str) > 0 {
		if r, strOffset = utf8.DecodeRuneInString(str); r != utf8.RuneError {
			switch {
			case r == '\\':
				if isEscaping {
					isEscaping = false
					sb.WriteRune(r)
					lastCh = r
				} else {
					isEscaping = true
				}
			case unicode.IsDigit(r):
				if isEscaping {
					isEscaping = false
					sb.WriteRune(r)
					lastCh = r
					break
				}
				if lastCh == utf8.RuneError {
					return "", NewErrLonelyQuantifier(r, pos)
				} else {
					retries, _ := strconv.ParseInt(string(r), 10, 32)
					if retries == 0 {
						return "", NewErrZeroQuantifier(r, pos)
					}
					// int(retries-1) because one instance of the lastCh is already in it
					for i := 0; i < int(retries-1); i++ {
						sb.WriteRune(lastCh)
					}
					lastCh = utf8.RuneError
				}
			case unicode.IsLetter(r):
				if isEscaping {
					return "", NewErrInvalidEscapeSeq(r, pos)
				}
				sb.WriteRune(r)
				lastCh = r
			default:
				return "", NewErrUnsupportedRune(r, pos)
			}
		} else {
			// because up to that moment `str` is a valid utf8 string
			panic("unreachable")
		}
		pos++
		str = str[strOffset:]
	}
	if isEscaping {
		return "", NewErrUnfinishedEscapeSeq(r, pos-1)
	}
	return sb.String(), nil
}
