package sort

type Config struct {
	// Action types
	JustCheck bool

	// Sort types
	SortKey          SortKey
	SortNumeric      bool
	SortMonth        bool
	SortHumanNumeric bool

	// Sort modifiers
	Reverse             bool
	IgnoreLeadingBlanks bool

	// Result modifiers
	UniqueOnly bool
}

type SortKey struct {
	Enabled bool
	KeyDef  string
}
