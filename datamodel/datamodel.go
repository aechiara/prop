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
package datamodel

import (
	"fmt"
	"io"
)

type ConfigLine struct {
	LineNo int
	Line   string
	Key    string
	Value  string
}

type ConfigFile struct {
	Lines []ConfigLine
}

var configFile *ConfigFile

func New() *ConfigFile {
	configFile = &ConfigFile{}
	return configFile
}

func (c *ConfigFile) WriteToFile(fo io.Writer) error {
	var lineToWrite string
	for _, k := range c.Lines {
		// log.Printf("[%d] - Key: [%s] => [%s]\n", k.LineNo, k.Key, k.Value)
		lineToWrite = k.Line
		if len(k.Value) > 0 {
			lineToWrite = fmt.Sprintf("%s = %s\n", k.Key, k.Value)
		}
		fo.Write([]byte(lineToWrite))
	}
	return nil
}

func (c *ConfigFile) ChangeValue(key string, value string) error {
	found := false
	cfLine := ConfigLine{}
	var idx int
	for i, r := range c.Lines {
		if r.Key == key {
			cfLine.Key = r.Key
			cfLine.Value = value
			cfLine.Line = r.Line
			cfLine.LineNo = r.LineNo
			found = true
			idx = i
			break
		}
	}

	if !found {
		return fmt.Errorf("Key %s not found", key)
	}

	temp := c.Lines[idx+1:]
	c.Lines = append(c.Lines[:idx], cfLine)
	c.Lines = append(c.Lines, temp...)

	return nil
}

func (c *ConfigFile) CheckValue(key string) (string, error) {
	found := false
	value := ""
	for _, s := range configFile.Lines {
		if s.Key == key {
			found = true
			value = s.Value
			break
		}
	}

	if !found {
		return "", fmt.Errorf("The key [%s] was not found", key)
	}

	return value, nil
}

func (c *ConfigFile) AddLine(ln ConfigLine) error {
	c.Lines = append(c.Lines, ln)

	return nil
}
