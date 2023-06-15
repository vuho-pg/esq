package esq

type ScoreMode string

const (
	ScoreModeAvg  ScoreMode = "avg"
	ScoreModeMax  ScoreMode = "max"
	ScoreModeMin  ScoreMode = "min"
	ScoreModeNone ScoreMode = "none"
	ScoreModeSum  ScoreMode = "sum"
)
