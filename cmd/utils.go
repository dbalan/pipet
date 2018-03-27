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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dbalan/pipet/pipetdata"
)

// console colors
var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()

func errorGuard(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", Red(msg), err)
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

func isValidDirectory(p string) bool {
	if strings.Contains(p, " ") {
		return false
	}

	fullPath := expandHome(p)
	fi, err := os.Stat(fullPath)
	if err == nil && !fi.IsDir() {
		return false // present and is a file
	}

	return true
}

func ensureConfig(cmd *cobra.Command, args []string) error {
	if _, ok := viper.Get("document_dir").(string); !ok {
		return errors.New("no document_dir set in config, run pipet init")
	}

	if _, ok := viper.Get("editor_binary").(string); !ok {
		return errors.New("no editor_binary set in conifg, run pipet init")
	}
	return nil
}

func getDataStore() *pipetdata.DataStore {
	diskPath := viper.Get("document_dir").(string)
	dataStore, err := pipetdata.NewDataStore(expandHome(diskPath))
	errorGuard(err, "error accessing data store")
	return dataStore
}

// call external editor to edit the snippet.
func editSnippet(fn string) error {
	editorBin := viper.Get("editor_binary").(string)
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

	coli := []string{}
	// compact the output
	for _, v := range oli {
		if v != "" {
			coli = append(coli, v)
		}
	}
	if len(coli) < 2 {
		return "", errors.New("bad data")
	}
	return coli[len(coli)-1], nil
}

func which(c string) (string, error) {
	var w bytes.Buffer

	cmd := exec.Command("which", c)
	cmd.Stdout = &w

	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(w.String(), "\n"), nil
}

// basic bare bones wrapper that calls fzf
// calls fzf on searchText and returns the selected line
func fuzzyWrapper(searchText string) (sid string, e error) {

	fzf, err := which("fzf")
	if err != nil {
		e = err
		return
	}

	var w bytes.Buffer

	cmd := exec.Command(fzf)

	cmd.Stdin = strings.NewReader(searchText)
	cmd.Stdout = &w
	cmd.Stderr = os.Stderr

	err = cmd.Start()
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

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	errorGuard(err, "reading failed")

	return strings.TrimSuffix(text, "\n")
}

func renderSnippetList(sns []*pipetdata.Snippet, header bool) string {
	output := []string{}
	if header {
		output = append(output, "Title | Tags | UID")
	}

	for _, snip := range sns {
		tags := strings.Join(snip.Meta.Tags, ",")
		if header {
			snip.Meta.Title = Green(snip.Meta.Title)
			tags = Blue(tags)
		}
		out := fmt.Sprintf("%s | %s | %s", snip.Meta.Title,
			tags, snip.Meta.UID)
		output = append(output, out)
	}

	return columnize.SimpleFormat(output)
}

func searchFullSnippet() (sid string, e error) {
	dataStore := getDataStore()

	sns, err := dataStore.List()
	if err != nil {
		e = errors.Wrap(err, "listing dataStore failed")
		return
	}

	rendered := renderSnippetList(sns, false)
	sid, err = fuzzyWrapper(rendered)
	if err != nil {
		e = errors.Wrap(err, "searching failed")
		return
	}
	return sid, err
}
