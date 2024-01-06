package unpack

import (
	"fmt"
)

type ErrLexicalBase struct {
	quantifier rune
	position   int
}

func (err ErrLexicalBase) Error() string {
	return fmt.Sprintf("%d:%q", err.position, err.quantifier)
}

func (err ErrLexicalBase) Is(target error) bool {
	e, ok := target.(ErrLexicalBase)
	return ok && e == err
}

type ErrLonelyQuantifier struct {
	ErrLexicalBase
}

func (err ErrLonelyQuantifier) Error() string {
	return fmt.Sprintf("lonely quantifier without a symbol in front of it: %s", err.ErrLexicalBase)
}

func (err ErrLonelyQuantifier) Is(target error) bool {
	e, ok := target.(ErrLonelyQuantifier)
	return ok && e == err
}

type ErrZeroQuantifier struct {
	ErrLexicalBase
}

func (err ErrZeroQuantifier) Error() string {
	return fmt.Sprintf("zero quantifier: %s", err.ErrLexicalBase)
}

func (err ErrZeroQuantifier) Is(target error) bool {
	e, ok := target.(ErrZeroQuantifier)
	return ok && e == err
}

type ErrUnsupportedRune struct {
	ErrLexicalBase
}

func (err ErrUnsupportedRune) Error() string {
	return fmt.Sprintf("unsupported rune: %s", err.ErrLexicalBase)
}

func (err ErrUnsupportedRune) Is(target error) bool {
	e, ok := target.(ErrUnsupportedRune)
	return ok && e == err
}

type ErrNotValidUtf8 struct{}

func (err ErrNotValidUtf8) Error() string {
	return "string is not a valid utf8"
}

func (err ErrNotValidUtf8) Is(target error) bool {
	_, ok := target.(ErrNotValidUtf8)
	return ok
}

type ErrUnfinishedEscapeSeq struct {
	ErrLexicalBase
}

func (err ErrUnfinishedEscapeSeq) Error() string {
	return fmt.Sprintf("unfinished escage seq: %s", err.ErrLexicalBase)
}

func (err ErrUnfinishedEscapeSeq) Is(target error) bool {
	e, ok := target.(ErrUnfinishedEscapeSeq)
	return ok && e == err
}

type ErrInvalidEscapeSeq struct {
	ErrLexicalBase
}

func (err ErrInvalidEscapeSeq) Error() string {
	return fmt.Sprintf("unfinished escape seq: %s", err.ErrLexicalBase)
}

func (err ErrInvalidEscapeSeq) Is(target error) bool {
	e, ok := target.(ErrInvalidEscapeSeq)
	return ok && e == err
}

func NewErrInvalidEscapeSeq(quantifier rune, position int) error {
	return ErrInvalidEscapeSeq{ErrLexicalBase{quantifier, position}}
}

func NewErrUnfinishedEscapeSeq(quantifier rune, position int) error {
	return ErrUnfinishedEscapeSeq{ErrLexicalBase{quantifier, position}}
}

func NewErrLonelyQuantifier(quantifier rune, position int) error {
	return ErrLonelyQuantifier{ErrLexicalBase{quantifier, position}}
}

func NewErrNotValidUtf8() error {
	return ErrNotValidUtf8{}
}
func NewErrZeroQuantifier(quantifier rune, position int) error {
	return ErrZeroQuantifier{ErrLexicalBase{quantifier, position}}
}

func NewErrUnsupportedRune(r rune, position int) error {
	return ErrUnsupportedRune{ErrLexicalBase{r, position}}
}
