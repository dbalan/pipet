package pipetdata

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var exampleSnippet = []byte(`---
uid: 1af02551-5821-4a54-be81-1a07441101f8
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
	assert.Equal(t, "1af02551-5821-4a54-be81-1a07441101f8", snip.Meta.UID, "uid should match")

	assert.Equal(t, "uname -a\n", snip.Data, "data should match")
}

func TestUnmarshal(t *testing.T) {
	snip := &Snippet{
		Meta: metadata{UID: "1af02551-5821-4a54-be81-1a07441101f8",
			Title: "Kernel version", Tags: []string{"linux", "kernel", "systems", "code"}},
		Data: "uname -a",
	}

	data, err := snip.Marshal()

	assert.Nil(t, err, "error should not have happened")
	assert.Equal(t, exampleSnippet, data, "render data should match")
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
	uid := ours.Name()

	assert.Equal(t, fn, filepath.Join(tmpdir, uid), "filepaths should match")
	// try to read back!
	sn, err := ds.Read(ours.Name())
	assert.Nil(t, err, "should be a valid snippet")

	expected := &Snippet{
		Meta: metadata{UID: uid, Title: "Kernel version", Tags: []string{"linux", "kernel", "systems", "code"}},
	}

	assert.Equal(t, expected.Meta, sn.Meta, "snippet metadata should match")

	// test delete
	err = ds.Delete(uid)
	assert.Nil(t, err, "should have deleted properly")

	// try random file
	err = ds.Delete("probably.txt")
	assert.NotNil(t, err, "should return no file")

	_, err = ds.Read(uid)
	assert.NotNil(t, err, "no such snippet should exist.")

}

func TestDataStoreList(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "testpipet")
	assert.Nil(t, err, "tmp directory")

	ds, err := NewDataStore(tmpdir)
	assert.Nil(t, err, "data store creation")

	snli, err := ds.List()
	assert.NotNil(t, err, "Should have errored")

	fn, err := ds.New("Kernel version", "linux", "kernel", "systems", "code")
	assert.Nil(t, err, "new snippet must be created")

	uid := filepath.Base(fn)
	// create 1

	snli, err = ds.List()
	assert.Nil(t, err, "should not error")
	assert.Len(t, snli, 1, "empty ds")

	expected := &Snippet{
		Meta: metadata{UID: uid, Title: "Kernel version", Tags: []string{"linux", "kernel", "systems", "code"}},
	}

	assert.Equal(t, expected.Meta, snli[0].Meta, "metadata must match")

	_, err = ds.New("Kernel version2", "linux", "kernel", "systems", "code")
	assert.Nil(t, err, "new snippet must be created")

	snli, err = ds.List()
	assert.Nil(t, err, "should not error")
	assert.Len(t, snli, 2, "empty ds")
}
