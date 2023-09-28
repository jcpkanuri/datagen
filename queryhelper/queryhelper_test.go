package queryhelper

import (
	"fmt"
	"log"
	"testing"

	"datagen/config"
	"datagen/generators"
)

func TestGetInsertStatements(t *testing.T) {

	got, err := config.ProbeTableMetadata("mysinglestore", "employee")
	got = config.Enrich("mysinglestore", got)
	got.Columns[1].Generator = generators.Generator{Name: "firstname"}

	got.UniqCols = make(map[string][]string)

	if err != nil {
		fmt.Println("Failed to load employee info")
	}

	result, err := GetInsertStatementPrefix(got)

	if err != nil {
		log.Fatal("Encountered error ", err)
	}

	result2, _ := GetRows(got, 10)

	fmt.Printf("Result1: %s", result)

	for i, v := range result2 {
		fmt.Printf("%d === %s\n", i, v.ValuesString())
	}

}

func TestGenerateFullInsertStatement(t *testing.T) {

	table, _ := config.ProbeTableMetadata("mysinglestore", "employee")
	table = config.Enrich("mysinglestore", table)
	table.Columns[1].Generator = generators.Generator{Name: "firstname"}

	table.UniqCols = make(map[string][]string)

	got, _ := GenerateFullInsertStatement(table, 10)

	fmt.Println(got)

}
