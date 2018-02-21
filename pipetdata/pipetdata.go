// Package pipetdata impliments meat of pipet, data access functions - which are
// being called out from CLI
package pipetdata

// DataStore is the main structure for snippet access
type DataStore struct {
	documentDir string
}

// Metadata for snippet
type metadata struct {
	Title string
	Tags  []string
}

// Snippet is the data type holding the actual snippet
type Snippet struct {
	meta metadata
	Data string
}

// MarshalMD serializes snippet data into markdown. Uses markdown with yaml
// metadata block on top
func (s *Snippet) MarshalMD() []byte {
	return []byte("NA")
}

// UnmarshalMD takes in markdown data with metadata blocks in yaml and populates
// Snippet structure.
func (s *Snippet) UnmarshalMD(data []byte) error {
	return nil
}
