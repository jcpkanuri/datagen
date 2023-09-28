package config

import (
	"os"
	"strings"
	"testing"
)

func TestWriteContent(t *testing.T) {

	content := "Disney Mickey Mouse"

	WriteContent("tmp_FileWrite.out", content)

	fcontent, err := os.ReadFile("tmp_FileWrite.out")

	if err != nil {
		t.Error("Unable to open file")
	}

	if strings.Compare(content, string(fcontent)) != 0 {
		t.Error("Content did not match")
	}

	t.Cleanup(func() {
		os.Remove("tmp_FileWrite.out")
	})
}

func TestWriteJson(t *testing.T) {

	empTable, err := ProbeTableMetadata("mysinglestore", "employee")
	if err != nil {
		t.Error("Failed to read table metadata")
	}

	WriteJsonFile(empTable, "tmp_employee.json")

	inf, err := os.Stat("tmp_employee.json")

	if err != nil {
		t.Error("File does not exists, Write Fail")
	}

	if inf.IsDir() || inf.Size() == 0 {
		t.Error("Empty File, write failed")
	}

	t.Cleanup(func() {
		os.Remove("tmp_employee.json")
	})
}

func TestWriteYaml(t *testing.T) {

	empTable, err := ProbeTableMetadata("mysinglestore", "employee")
	if err != nil {
		t.Error("Failed to read table metadata")
	}

	WriteYamlFile(empTable, "tmp_employee.yaml")

	inf, err := os.Stat("tmp_employee.yaml")

	if err != nil {
		t.Error("File does not exists, Write Fail")
	}

	if inf.IsDir() || inf.Size() == 0 {
		t.Error("Empty File, write failed")
	}

	t.Cleanup(func() {
		os.Remove("tmp_employee.yaml")
	})
}
