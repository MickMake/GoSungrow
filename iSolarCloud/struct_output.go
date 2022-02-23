package iSolarCloud

const (
	TypeNone  = iota
	TypeJson  = iota
	TypeFile  = iota
	TypeRaw   = iota
	TypeHuman = iota

	StringTypeNone  = ""
	StringTypeJson  = "json"
	StringTypeFile  = "file"
	StringTypeRaw   = "raw"
	StringTypeHuman = "human"
)

type OutputType int

func (out *OutputType) SetNone() {
	*out = TypeNone
}
func (out *OutputType) SetJson() {
	*out = TypeJson
}
func (out *OutputType) SetFile() {
	*out = TypeFile
}
func (out *OutputType) SetRaw() {
	*out = TypeRaw
}
func (out *OutputType) SetHuman() {
	*out = TypeHuman
}

func (out *OutputType) IsNone() bool {
	if *out == TypeNone {
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
func (out *OutputType) IsFile() bool {
	if *out == TypeFile {
		return true
	}
	return false
}
func (out *OutputType) IsRaw() bool {
	if *out == TypeRaw {
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

func (out *OutputType) IsStrNone(t string) bool {
	if t == StringTypeNone {
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
func (out *OutputType) IsStrFile(t string) bool {
	if t == StringTypeFile {
		return true
	}
	return false
}
func (out *OutputType) IsStrRaw(t string) bool {
	if t == StringTypeRaw {
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
