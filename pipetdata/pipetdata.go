// Package pipetdata impliments meat of pipet, data access functions - which are
// being called out from CLI
package pipetdata

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	EBadData = fmt.Errorf("bad data")
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

func splitData(buf []byte) (front, data []byte, err error) {
	guard := []byte("---")
	lg := len(guard)

	if !bytes.HasPrefix(buf, guard) {
		err = EBadData
		return
	}

	buf = buf[lg:]

	end := bytes.Index(buf, guard)
	if end == -1 {
		err = EBadData
		return
	}

	front = buf[0:end]
	// skip --- + '\n'
	data = buf[end+lg+1:]
	return
}

// Unmarshal takes in note data with metadata blocks in yaml and populates
// Snippet structure.
// https://jekyllrb.com/docs/frontmatter/
func (s *Snippet) Unmarshal(buf []byte) error {
	front, data, err := splitData(buf)

	var meta metadata
	yaml.Unmarshal(front, &meta)
	s.Meta = meta
	s.Data = string(data)
	return err
}
