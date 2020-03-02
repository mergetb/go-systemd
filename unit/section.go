package unit

import ()

// UnitEntry is a single line entry in a Unit file.
type UnitEntry struct {
	Name  string
	Value string
}

// UnitSection is a section in a Unit file. The section name
// and a list of entries in that section.
type UnitSection struct {
	Section string
	Entries []*UnitEntry
}

// String implements the stringify interface for UnitEntry
func (u *UnitEntry) String() string {
	return "{Name: " + u.Name + ", " + "Value: " + u.Value + "}"
}

// String implements the stringify interface for UnitSection
func (s *UnitSection) String() string {
	result := "{Section: " + s.Section
	for _, e := range s.Entries {
		result += e.String()
	}

	result += "}"
	return result
}
