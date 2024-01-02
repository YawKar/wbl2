package sort

type Config struct {
	SortKey          SortKey
	SortNumeric      bool
	SortReverse      bool
	SortMonth        bool
	SortHumanNumeric bool

	JustCheck bool

	IgnoreLeadingBlanks bool

	UniqueOnly bool
}

type SortKey struct {
	Enabled bool
	KeyDef  string
}
