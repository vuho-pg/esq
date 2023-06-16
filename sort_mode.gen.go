package esq

type SortMode string

const (
	SortModeMin    SortMode = "min"
	SortModeMax    SortMode = "max"
	SortModeSum    SortMode = "sum"
	SortModeAvg    SortMode = "avg"
	SortModeMedian SortMode = "median"
)
