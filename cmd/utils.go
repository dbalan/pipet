// Copyright © 2018 Dhananjay Balan
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dbalan/pipet/pipetdata"

	"bytes"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func errorGuard(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
		os.Exit(-1)
	}
}

func expandHome(p string) string {
	if !strings.HasPrefix(p, "~/") {
		return p
	}

	h, err := homedir.Dir()
	if err != nil {
		return p
	}

	return filepath.Join(h, p[2:])
}

func getDataStore() *pipetdata.DataStore {
	diskPath := viper.Get("document_dir").(string)
	dataStore, err := pipetdata.NewDataStore(expandHome(diskPath))
	errorGuard(err, "error accessing data store")
	return dataStore
}

// call external editor to edit the snippet.
func editSnippet(fn string) error {
	editorBin := os.Getenv("EDITOR")
	if editorBin == "" {
		return errors.New("please populate default EDITOR value by setting $EDITOR")
	}

	cmd := exec.Command(editorBin, fn)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return errors.Wrap(err, "launching editer failed")
	}
	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err, "editing failed")
	}
	return nil
}

func parseOutput(out string) (string, error) {
	out = strings.TrimSuffix(out, "\n")
	oli := strings.Split(out, " ")
	if len(oli) < 2 {
		return "", errors.New("bad data")
	}
	return oli[0], nil
}

// basic bare bones wrapper that calls fzf
// calls fzf on searchText and returns the selected line
func fuzzyWrapper(searchText string) (sid string, e error) {
	// checks in initconfig ensures that this doesn't panic. Its not great,
	// but its fine.
	fzy := viper.Get("fzf").(string)

	var w bytes.Buffer

	cmd := exec.Command(fzy)

	cmd.Stdin = strings.NewReader(searchText)
	cmd.Stdout = &w
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		e = errors.Wrap(err, "launching editer failed")
		return
	}

	err = cmd.Wait()
	if err != nil {
		e = errors.Wrap(err, "editing failed")
		return
	}
	return parseOutput(w.String())
}
