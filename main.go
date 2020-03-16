package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aechiara/edprop/datamodel"
	"github.com/aechiara/edprop/utils"
)

func main() {
	log.Println("Iniciando editor de properties")

	configFile := datamodel.New()

	// TODO: parametro via linha de comando
	err := readConfig("LocalConfig.txt", configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Change Value
	err = configFile.ChangeValue("tskJarPath", "/tmp/sleuthkit/sleuth.jar")
	if err != nil {
		log.Fatalln("Error changing value", err)
	}

	fo, err := os.Create("saida.txt")
	if err != nil {
		log.Fatalf("Unable to create output file: %v", err)
	}
	defer fo.Close()

	err = configFile.WriteToFile(fo)

	log.Println("Done")

}

func readConfig(configFile string, confStruct *datamodel.ConfigFile) error {
	if !utils.FileExists(configFile) {
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
