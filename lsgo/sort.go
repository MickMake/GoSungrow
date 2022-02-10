package lsgo

// BySize tells `sort.Sort` how to sort by file size
type BySize []*DisplayItem

func (s BySize) Less(i, j int) bool {
	return s[i].info.Size() < s[j].info.Size()
}
func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ByTime tells `sort.Sort` how to sort by last modified time
type ByTime []*DisplayItem

func (s ByTime) Less(i, j int) bool {
	return s[i].info.ModTime().Unix() < s[j].info.ModTime().Unix()
}
func (s ByTime) Len() int {
	return len(s)
}
func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ByKind tells `sort.Sort` how to sort by file extension
type ByKind []*DisplayItem

func (s ByKind) Less(i, j int) bool {
	var kindi, kindj string
	if s[i].basename == "" {
		kindi = "."
	} else if s[i].ext == "" {
		kindi = "0"
	} else {
		kindi = s[i].ext
	}
	if s[j].basename == "" {
		kindj = "."
	} else if s[j].ext == "" {
		kindj = "0"
	} else {
		kindj = s[j].ext
	}
	if kindi == kindj {
		if kindi == "." {
			return s[i].ext < s[j].ext
		}
		return s[i].basename < s[j].basename
	}
	return kindi < kindj
}
func (s ByKind) Len() int {
	return len(s)
}
func (s ByKind) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func reverse(s []*DisplayItem) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
