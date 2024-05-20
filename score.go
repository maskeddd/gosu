package gosu

import "time"

type Score struct {
	ID         int             `json:"id"`
	UserID     int             `json:"user_id"`
	Accuracy   float64         `json:"accuracy"`
	Mods       []string        `json:"mods"`
	Score      int             `json:"score"`
	MaxCombo   int             `json:"max_combo"`
	Perfect    bool            `json:"perfect"`
	Statistics ScoreStatistics `json:"statistics"`
	Passed     bool            `json:"passed"`
	Pp         float32         `json:"pp"`
	Rank       Grade           `json:"rank"`
	CreatedAt  time.Time       `json:"created_at"`
	Mode       Ruleset         `json:"mode"`
	ModeInt    int             `json:"mode_int"`
	Replay     bool            `json:"replay"`
}

type ScoreStatistics struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}

type ScoreWeight struct {
	Percentage float32 `json:"percentage"`
	PP         float32 `json:"pp"`
}
