// Package pipetdata impliments meat of pipet, data access functions - which are
// being called out from CLI
package pipetdata

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	// EBadData error bad data on disk
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

// NewDataStore creates a new datastore abastraction for storing notes. disk
// path is passed as documentDir
func NewDataStore(documentDir string) (*DataStore, error) {
	fi, err := os.Stat(documentDir)
	if err != nil {
		//create directory
		err = os.MkdirAll(documentDir, 0644)
		if err != nil {
			return nil, err
		}
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("path is not a directory: %s", documentDir)
	}

	return &DataStore{documentDir}, nil
}

func snippetStoreName(title string) string {
	unusable := []string{" ", "/", "*", "%"}
	for _, un := range unusable {
		title = strings.Replace(title, un, "-", -1)
	}
	return fmt.Sprintf("%s.txt", title)
}

// Exist checks with a snippet with the name exists.
func (d *DataStore) Exist(filename string) bool {
	fullpath := filepath.Join(d.documentDir, filename)
	_, err := os.Stat(fullpath)
	return err == nil
}

// New creates a new entry in snippets
func (d *DataStore) New(title string, tags ...string) (fn string, err error) {
	ns := &Snippet{
		Meta: metadata{title, tags},
	}

	snipname := snippetStoreName(title)
	if d.Exist(snipname) {
		return "", errors.New("duplicate snippet")
	}

	filename := filepath.Join(d.documentDir, snipname)
	data, err := ns.Marshal()
	if err != nil {
		return "", errors.Wrap(err, "marshalling failed")
	}

	err = ioutil.WriteFile(filename, data, 0755)
	return filename, err
}