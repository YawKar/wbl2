package sort

type Config struct {
	OrderingOptions
	OtherOptions
}

type OrderingOptions struct {
	SortKey          *SortKey
	SortNumeric      bool
	SortMonth        bool
	SortHumanNumeric bool

	Reverse             bool
	IgnoreLeadingBlanks bool
}

type OtherOptions struct {
	UniqueOnly bool
}

type SortKey struct {
	StartField int
	StartChar  int
	StartOpts  *OrderingOptions
	EndField   int
	EndChar    int
}
