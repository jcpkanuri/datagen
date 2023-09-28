package config

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"datagen/types"
)

func TestProbeTableMetadata(t *testing.T) {

	got, err := ProbeTableMetadata("mysinglestore", "employee")

	if err != nil {
		t.Errorf("Failed to parse table info")
	}

	if strings.Compare(got.Name, "employee") != 0 {
		t.Error("Failed to fetch right table")
	}

	if strings.Compare(got.Columns[1].Name, "ename") != 0 {
		t.Error("Failed to fetch column name")
	}

	if !strings.Contains(got.Columns[0].Extra, "increment") {
		t.Error("Failed to fetch extra field info", got.Columns[0].Extra)
	}
}

func TestParseJson(t *testing.T) {

	empTable, err := ProbeTableMetadata("mysinglestore", "employee")
	if err != nil {
		t.Error("Failed to read table metadata")
	}

	WriteJsonFile(empTable, "tmp_employee.json")

	empTable2, err := ParseJsonFile("tmp_employee.json")

	if err != nil {
		t.Error("Failed to read json file")
	}

	if !reflect.DeepEqual(empTable, empTable2) {
		t.Error("Could not parse json well")
	}

	// t.Cleanup(func() {
	// 	os.Remove("tmp_employee.json")
	// })
}

func TestParseYaml(t *testing.T) {

	empTable, err := ProbeTableMetadata("mysinglestore", "employee")
	if err != nil {
		t.Error("Failed to read table metadata")
	}

	WriteYamlFile(empTable, "tmp_employee.yaml")

	empTable2, err := ParseYamlFile("tmp_employee.yaml")

	if err != nil {
		t.Error("Failed to read yaml file")
	}

	if empTable.Name != empTable2.Name {
		fmt.Println("Names is problem")
	}

	// if !Equal(empTable.Columns, empTable2.Columns) {
	// 	fmt.Println("Cols are prob")
	// }

	// t.Cleanup(func() {
	// 	os.Remove("tmp_employee.yaml")
	// })
}

func Equal(a, b []types.Column) bool {
	if len(a) != len(b) {
		fmt.Println("length not matching")
		return false
	}

	//todo compare struts fields
	for i, v := range a {
		if v.Name != b[i].Name {
			fmt.Printf("index %d not matching", i)
			return false
		}
	}
	return true
}
