package config

import (
	"encoding/json"
	"log"
	"os"

	"datagen/types"

	"gopkg.in/yaml.v3"
)

func WriteContent(outFileName string, fileContent string) {

	f, err := os.Create(outFileName)

	if err != nil {
		log.Fatal("Failed to create file", outFileName, err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fileContent)

	if err2 != nil {
		log.Fatal("Failed to write content", outFileName, err2)
	}

}

func WriteJsonFile(t types.Table, outFile string) {

	jsonContent, _ := json.Marshal(t)
	result, _ := PrettyString(string(jsonContent))

	WriteContent(outFile, result)
}

func WriteYamlFile(t types.Table, outFile string) {
	yamlContent, _ := yaml.Marshal(t)
	WriteContent(outFile, string(yamlContent))
}
