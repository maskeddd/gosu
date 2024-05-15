package gosu

// ENUM(osu, taiko, fruits, mania)
type Ruleset int

// ENUM(graveyard = -2, WIP, pending, ranked, approved, qualified, loved)
type RankStatus int

type Grade string

const (
	GradeSS  Grade = "SS"
	GradeSSH Grade = "SSH"
	GradeS   Grade = "S"
	GradeSH  Grade = "SH"
	GradeA   Grade = "A"
	GradeB   Grade = "B"
	GradeC   Grade = "C"
	GradeD   Grade = "D"
	GradeF   Grade = "F"
)

type Cursor interface{}

type Sort string

const (
	SortDescending = "id_desc"
	SortAscending  = "id_asc"
)
