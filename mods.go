package gosu

import "strings"

type Mod int

const (
	NF Mod = 1 << iota
	EZ
	TD
	HD
	HR
	SD
	DT
	RX
	HT
	NC
	FL
	AT
	SO
	AP
	PF
	K4
	K5
	K6
	K7
	K8
	FI
	RN
	CN
	TR
	K9
	KC
	K1
	K3
	K2
	SV2
	MR
)

var modStrings = map[Mod]string{
	NF:  "NF",
	EZ:  "EZ",
	TD:  "TD",
	HD:  "HD",
	HR:  "HR",
	SD:  "SD",
	DT:  "DT",
	RX:  "RX",
	HT:  "HT",
	NC:  "NC",
	FL:  "FL",
	AT:  "AT",
	SO:  "SO",
	AP:  "AP",
	PF:  "PF",
	K4:  "K4",
	K5:  "K5",
	K6:  "K6",
	K7:  "K7",
	K8:  "K8",
	FI:  "FI",
	RN:  "RN",
	CN:  "CN",
	TR:  "TR",
	K9:  "K9",
	KC:  "KC",
	K1:  "K1",
	K3:  "K3",
	K2:  "K2",
	SV2: "SV2",
	MR:  "MR",
}

func ModsToString(mods Mod) string {
	var result strings.Builder

	for i := NF; i <= MR; i <<= 1 {
		if mods&i != 0 {
			result.WriteString(modStrings[i])
		}
	}

	return result.String()
}

func ModsToStrings(mods Mod) []string {
	var result []string

	for i := NF; i <= MR; i <<= 1 {
		if mods&i != 0 {
			result = append(result, modStrings[i])
		}
	}

	return result
}
