// Copyright © 2018 Dhananjay Balan <mail@dbalan.in>
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

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [uid]",
	Short:   "Remove snippet from storage (this is irreversible!)",
	Args:    cobra.MaximumNArgs(1),
	PreRunE: ensureConfig,
	Run: func(cmd *cobra.Command, args []string) {
		sid := ""
		if len(args) == 0 {
			s, err := searchFullSnippet()
			errorGuard(err, "")
			sid = s
		} else {
			sid = args[0]
		}

		dataStore := getDataStore()
		snip, err := dataStore.Read(sid)
		errorGuard(err, "querying snippet failed")

		fmt.Printf("Are your you want to %s '%s' [y/n]: ", Red("DELETE"), Green(snip.Meta.Title))
		confirm := readLine()
		if confirm == "y" || confirm == "yes" {
			errorGuard(dataStore.Delete(sid), "program failed to delete")
			fmt.Println("deleted!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// FIXME: add an archive
}
