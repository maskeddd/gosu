package gosu

// ENUM(osu, taiko, fruits, mania)
type Ruleset int

// ENUM(graveyard = -2, WIP, pending, ranked, approved, qualified, loved)
type RankStatus int

type Grade string

const (
	SS  Grade = "SS"
	SSH Grade = "SSH"
	S   Grade = "S"
	SH  Grade = "SH"
	A   Grade = "A"
	B   Grade = "B"
	C   Grade = "C"
	D   Grade = "D"
	F   Grade = "F"
)

type Cursor map[string]interface{}
