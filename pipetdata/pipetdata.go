// Package pipetdata impliments meat of pipet, data access functions - which are
// being called out from CLI
package pipetdata

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// DataStore is the main structure for snippet access
type DataStore struct {
	documentDir string
}

// Metadata for snippet
type metadata struct {
	Title string   `yaml:"title"`
	Tags  []string `yaml:"tags,omitempty"`
}

// Snippet is the data type holding the actual snippet
type Snippet struct {
	Meta metadata
	Data string
}

// Marshal serializes snippet data into bytes. Format is
// ---
// yaml metadata front
// ---
// <text follows>
// This is very similiar to pandoc markdown except its just arbitary text for now.
func (s *Snippet) Marshal() ([]byte, error) {
	template := `---
%s---
%s
`
	meta, err := yaml.Marshal(s.Meta)
	if err != nil {
		return []byte{}, errors.Wrap(err, "yaml rendering failed")
	}

	rendered := fmt.Sprintf(template, meta, s.Data)
	return []byte(rendered), nil
}

// UnmarshalMD takes in note data with metadata blocks in yaml and populates
// Snippet structure.
func (s *Snippet) Unmarshal(data []byte) error {
	return nil
}
