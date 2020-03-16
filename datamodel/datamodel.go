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

func (c *ConfigFile) AddLine(ln ConfigLine) error {
	c.Lines = append(c.Lines, ln)

	return nil
}
