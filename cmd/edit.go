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
	"os"
	"strings"

	"github.com/aechiara/prop/datamodel"
	"github.com/aechiara/prop/utils"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit config_file_path key_to_find:new_value",
	Short: "Edit a configuration on a properties file",
	Long: `Edit a configuration on a properties file, if the property does not exists, returns an error.
	No new properties are created, for that you should use the add command.`,
	Args:                  cobra.MinimumNArgs(2),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		props := strings.Split(args[1], ":")
		if len(args) > 2 || len(props) != 2 || len(props[1]) == 0 || len(props[0]) == 0 {
			return fmt.Errorf("Wrong parameter format, the format shoud be key:value")
		}

		// fmt.Printf("[%s] - [%s]:[%s]\n", args[0], props[0], props[1])

		fileName := args[0]

		configFile := datamodel.New()

		err := utils.ReadConfig(fileName, configFile)
		if err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}

		err = configFile.ChangeValue(props[0], props[1])
		if err != nil {
			return fmt.Errorf("Key [%s] was not found on file [%s]\n", args[0], fileName)
		}

		fo, err := os.Create("saida.txt")
		if err != nil {
			return fmt.Errorf("Unable to create output file: %v", err)
		}
		defer fo.Close()

		err = configFile.Write(fo)

		return err
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
