/*
Copyright © 2020 Alex Eduardo Chiaranda <aechiara@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/aechiara/prop/datamodel"
	"github.com/aechiara/prop/utils"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:                   "check config_file_path key_to_find",
	Short:                 "Check for a given parameter on a properties file",
	Args:                  cobra.MinimumNArgs(2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		fileName := args[0]

		configFile := datamodel.New()

		err := utils.ReadConfig(fileName, configFile)
		if err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}

		value, err := configFile.CheckValue(args[1])
		if err != nil {
			return fmt.Errorf("Key [%s] was not found on file [%s]\n", args[1], fileName)
		}
		fmt.Printf("Key [%s] value is [%s]\n", args[1], value)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
