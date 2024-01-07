package guts

type Config struct {
	FieldsRanges  []Range
	Delimiter     string
	OnlyDelimited bool
}

type Range struct {
	Leftx  int
	Rigthx int
}
