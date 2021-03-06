// Package cmd is the root for command line options
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add config_file_path key_to_add:new_value",
	Short: "Adds a new property to a properties file",
	Long: `Adds a new property to an existing properties file.
If the property already exists, the value WILL NOT be changed and an error will be returned,
unless a [ -f|--force ] flag is passed.`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		isForced, err := cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("Could not get Flag")
		}

		backup, err := cmd.Flags().GetBool("backup")
		if err != nil {
			return fmt.Errorf("Could not get Flag")
		}

		filename := args[0]
		props := strings.Split(args[1], ":")
		if len(args) > 2 || len(props) != 2 || len(props[1]) == 0 || len(props[0]) == 0 {
			return fmt.Errorf("Wrong parameter format, the format shoud be key:value")
		}

		// fmt.Printf("isForced: %v\n", isForced)
		// fmt.Printf("file:[%s] - key: [%s][%s]\n", fileName, props[0], props[1])

		configFile := datamodel.New()

		err = utils.ReadConfig(filename, configFile)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}

		if err != nil {
			return fmt.Errorf("error inserting new Property: %v", err)
		}

		_, err = configFile.CheckValue(props[0])
		// if key was FOUND and isForced not set
		if err == nil && !isForced {
			return fmt.Errorf("key [%s] already exists on file [%s], use --force to overwrite the property", props[0], filename)
		}

		confLine := datamodel.ConfigLine{
			Line:   strings.Join(props, " = "),
			LineNo: len(configFile.Lines),
			Key:    props[0],
			Value:  props[1],
		}

		// if key was NOT FOUND
		if err == nil && isForced {
			err = configFile.ChangeValue(props[0], props[1])
			fmt.Printf("Key [%s] FORCED CHANGE in [%s]\n", props[0], filename)
		} else {
			err = configFile.AddLine(confLine)
			fmt.Printf("Key [%s] added in [%s]\n", props[0], filename)
		}

		if backup {
			backupFilename := utils.GenBackupName(filename)
			err = os.Rename(filename, backupFilename)
			if err != nil {
				return fmt.Errorf("Unable to Create Backup File: %v", err)
			}
		}

		fo, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("Unable to create output file: %v", err)
		}
		defer fo.Close()

		err = configFile.Write(fo)

		return err
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("force", "f", false, "If the key already exists, replace with new value")
	addCmd.Flags().BoolP("backup", "b", false, "Creates a backup from original file")
}
