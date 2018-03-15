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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure pipet",
	Long:  `Creates the config files for pipet if not present, usually you only need at the first use`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile := expandHome("~/.pipet.yaml")
		_, err := os.Stat(configFile)
		if err == nil {
			errorGuard(fmt.Errorf("e"), "files exists: delete ~/.pipet.yaml if you want to reconfig")
		}

		fmt.Printf("Where do you want to store snippets? (directory|default: ~/snippets): ")
		snipDir := readLine()

		if !isValidDirectory(snipDir) {
			errorGuard(fmt.Errorf("invalid path"), "error")
		}

		// try getting $EDITOR!
		defEdit := "(no defaults)"
		editor := os.Getenv("EDITOR")
		if editor != "" {
			defEdit = fmt.Sprintf("(defaults to %s)", editor)
		}

		fmt.Printf("path to editor to use with pippet %s:", defEdit)
		eBin := readLine()

		config := &struct {
			DocDir string `yaml:"document_dir,omitempty"`
			EBin   string `yaml:"editor_binary,omitempty"`
		}{}

		if eBin == "" && defEdit == "" {
			errorGuard(fmt.Errorf("empy input"), "no editor specified")
		}

		if eBin == "" {
			eBin = editor
		}

		if eBin != "" {
			path, err := which(eBin)
			errorGuard(err, "No such editor found!")

			// expanded path
			if path != eBin {
				fmt.Printf("Using %s as the absoulte path to editor\n", Green(path))
			}
			config.EBin = path
		} else {
			errorGuard(fmt.Errorf("empty input"), "no editor sepcified")
		}

		if snipDir != "" {
			config.DocDir = snipDir
		} else {
			config.DocDir = expandHome("~/snippets")
		}

		buf, err := yaml.Marshal(config)

		if err != nil {
			errorGuard(err, "failed marshalling data")
		}
		errorGuard(ioutil.WriteFile(configFile, buf, 0755), "writing file failed")
		fmt.Printf(`pipet is now ready to use
 snippets are stored in: %s
 config is stored in %s
`, Green(snipDir), Green(configFile))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
