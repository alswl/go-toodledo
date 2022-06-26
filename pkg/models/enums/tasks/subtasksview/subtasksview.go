//go:generate stringer -type=Mode
package subtasksview

type Mode int

const (
	Inline   Mode = 0
	Hidden   Mode = 1
	Indented Mode = 2
)

var ModeAll = []Mode{
	Inline,
	Hidden,
	Indented,
}

var ModeMap = map[string]Mode{
	"inline":   Inline,
	"hidden":   Hidden,
	"indented": Indented,
}

func ModeValue2Type(input int64) Mode {
	for _, x := range ModeAll {
		if x == Mode(input) {
			return x
		}
	}
	return Inline
}

func ModeString2Type(input string) Mode {
	for k, v := range ModeMap {
		if k == input {
			return v
		}
	}
	return Inline
}
