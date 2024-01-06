package grib

type Config struct {
	SelectionConfig
	OutputConfig
	ContextConfig
	MiscellaneousConfig
}

type SelectionConfig struct {
	IgnoreCase   bool
	FixedStrings bool

	Patterns []string
}

type OutputConfig struct {
	LineNumber bool
	Count      bool
}

type ContextConfig struct {
	Before  int
	After   int
	Context int
}

type MiscellaneousConfig struct {
	InvertMatch bool
	Скрепы      bool
}
