package iSolarCloud

import "GoSungro/iSolarCloud/sungro"

var SunGro sungro.SunGro

// type Areas struct {
// 	domain.Domain
// }

const (
	TypeGit    = iota
	TypeJson   = iota
	TypeHuman  = iota
	TypeGoogle = iota

	StringTypeGit = "git"
	StringTypeJson = "json"
	StringTypeHuman = "human"
	StringTypeGoogle = "google"
)



type OutputType int
func (out *OutputType) IsGit() bool {
	if *out == TypeGit {
		return true
	}
	return false
}
func (out *OutputType) IsJson() bool {
	if *out == TypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsHuman() bool {
	if *out == TypeHuman {
		return true
	}
	return false
}
func (out *OutputType) IsGoogle() bool {
	if *out == TypeGoogle {
		return true
	}
	return false
}

func (out *OutputType) IsStrGit(t string) bool {
	if t == StringTypeGit {
		return true
	}
	return false
}
func (out *OutputType) IsStrJson(t string) bool {
	if t == StringTypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsStrHuman(t string) bool {
	if t == StringTypeHuman {
		return true
	}
	return false
}
func (out *OutputType) IsStrGoogle(t string) bool {
	if t == StringTypeGoogle {
		return true
	}
	return false
}
