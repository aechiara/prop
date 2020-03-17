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
	"io"
	"os"
	"strings"

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
	var line string
	totalLinhas := 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		totalLinhas++

		var key, value string
		b := []byte(line)
		// tratar # e nova linha e caracteres especiais no inicio dos arquivos
		// ef bb bf 23 23 23
		if len(line) > 2 && b[0] != 239 && !strings.HasPrefix(line, `#`) {
			// log.Print(line)
			splitted := strings.Split(line, "=")
			key, value = strings.TrimSpace(splitted[0]), strings.TrimSpace(splitted[1])
			// log.Printf("key: [%s] - value: [%s]", key, value)
		}

		configLine := datamodel.ConfigLine{
			LineNo: totalLinhas,
			Line:   line,
			Key:    key,
			Value:  value,
		}

		confStruct.AddLine(configLine)

		// fmt.Printf(" > Read %d characters\n", len(line))
		// fmt.Printf("line: [%s]\n", line)

	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	// log.Printf("Total de linhas: [%d]\n", totalLinhas)
	// log.Printf("Total de linhas: [%d]\n", len(confStruct.Lines))

	return nil
}
