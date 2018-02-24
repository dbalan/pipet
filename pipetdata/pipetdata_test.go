package pipetdata

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var exampleSnippet = []byte(`---
title: Kernel version
tags:
- linux
- kernel
- systems
- code
---
uname -a
`)

func TestMarshal(t *testing.T) {
	var snip Snippet
	err := snip.Unmarshal(exampleSnippet)
	assert.Nil(t, err, "error should not have happened")

	assert.Equal(t, "Kernel version", snip.Meta.Title, "title should match")
	assert.Equal(t, []string{"linux", "kernel", "systems", "code"},
		snip.Meta.Tags, "tags should match")

	assert.Equal(t, "uname -a\n", snip.Data, "data should match")
}

func TestUnmarshal(t *testing.T) {
	snip := &Snippet{
		Meta: metadata{"Kernel version", []string{"linux", "kernel", "systems", "code"}},
		Data: "uname -a",
	}

	data, err := snip.Marshal()

	assert.Nil(t, err, "error should not have happened")
	assert.Equal(t, exampleSnippet, data, "render data should match")
}

func TestSnippetFile(t *testing.T) {
	cases := []struct {
		Title string
		Fname string
	}{
		{"hello world", "hello-world.txt"},
		{"hello*world%inlive/what", "hello-world-inlive-what.txt"},
	}

	for _, c := range cases {
		assert.Equal(t, c.Fname, snippetStoreName(c.Title), "should match!")
	}
}

func TestDataStore(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "testpipet")
	assert.Nil(t, err, "tmp directory")

	ds, err := NewDataStore(tmpdir)
	assert.Nil(t, err, "data store creation")

	fn, err := ds.New("Kernel version", "linux", "kernel", "systems", "code")
	assert.Nil(t, err, "new snippet must be created")

	// try to read back file
	fi, err := ioutil.ReadDir(tmpdir)

	assert.Len(t, fi, 1, "should be just one file")
	ours := fi[0]
	assert.Equal(t, fn, filepath.Join(tmpdir, ours.Name()), "filename should match")

	// try to read back!
	sn, err := ds.Read(ours.Name())
	assert.Nil(t, err, "should be a valid snippet")

	expected := &Snippet{
		Meta: metadata{"Kernel version", []string{"linux", "kernel", "systems", "code"}},
	}

	assert.Equal(t, expected.Meta, sn.Meta, "snippet metadata should match")
}

func TestDataStoreList(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "testpipet")
	assert.Nil(t, err, "tmp directory")

	ds, err := NewDataStore(tmpdir)
	assert.Nil(t, err, "data store creation")

	snli, err := ds.List()
	assert.Nil(t, err, "should not error on empty ds")
	assert.Len(t, snli, 0, "empty ds")

	_, err = ds.New("Kernel version", "linux", "kernel", "systems", "code")
	assert.Nil(t, err, "new snippet must be created")

	expected := &Snippet{
		Meta: metadata{"Kernel version", []string{"linux", "kernel", "systems", "code"}},
	}

	snli, err = ds.List()
	assert.Nil(t, err, "should not error")
	assert.Len(t, snli, 1, "empty ds")
	assert.Equal(t, expected.Meta, snli[0].Meta, "metadata must match")

	_, err = ds.New("Kernel version2", "linux", "kernel", "systems", "code")
	assert.Nil(t, err, "new snippet must be created")

	snli, err = ds.List()
	assert.Nil(t, err, "should not error")
	assert.Len(t, snli, 2, "empty ds")
}
