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
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	snippetTags *[]string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Short:   `Creates a new snippet and opens editor to edit content`,
	PreRunE: ensureConfig,
	Run: func(cmd *cobra.Command, args []string) {
		title := cmd.Flag("title").Value.String()

		if title == "untitled" {
			fmt.Printf("Title for new snippet: ")
			if t := readLine(); t != "" {
				title = t
			}
		}

		dataStore := getDataStore()

		fn, err := dataStore.New(title, *snippetTags...)
		errorGuard(err, "creating snippet failed")

		err = editSnippet(fn)
		errorGuard(err, "opening snippet editor failed")

		fmt.Println("created a new snippet: ", fn)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newCmd.PersistentFlags().String("title", "untitled", "title for the snippet, if unset a random uuid is used.")
	snippetTags = newCmd.PersistentFlags().StringArray("tags", []string{"untagged"}, "tags for snippet, if unset, a single tag `untagged` is set.")
}
