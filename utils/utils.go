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
package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aechiara/prop/datamodel"
)

// FileExists check if path exists
func FileExists(path string) bool {
	exists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exists = false
	}

	return exists
}

// ReadConfig and loads it's content into a struct
func ReadConfig(configFile string, confStruct *datamodel.ConfigFile) error {
	if !FileExists(configFile) {
		return fmt.Errorf("%s does not exists", configFile)
	}

	file, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("Error trying to Open file [%s]: %v", configFile, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	return confStruct.Read(*reader)
}
