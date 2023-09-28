package types

import (
	"datagen/generators"
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {

	testCases := []struct {
		name  string
		input generators.Generator
	}{
		{"string", generators.Generator{Name: "string", Len: 10}},
		{"lexify", generators.Generator{Name: "lexify", Expression: "???"}},
		{"numerify", generators.Generator{Name: "numerify", Expression: "###"}},
		{"bothify", generators.Generator{Name: "bothify", Expression: "?? - ###"}},
		{"firstname", generators.Generator{Name: "firstname"}},
		{"lastname", generators.Generator{Name: "lastname"}},
	}

	for _, tc := range testCases {

		result := tc.input.Generate()

		if result == "" {
			t.Errorf("Generator %s failed", tc.input.Name)
		} else {
			fmt.Printf("Test Name: %s, Generated: %s \n", tc.name, tc.input.Generate())
		}
	}
}

func TestColumnValueGeneration(t *testing.T) {
	tinyCol := &Column{Name: "tinycolummn", DataType: "tinyint(2)"}

	fmt.Println("tinycol", tinyCol.GenerateNewValue())
}

func TestBigIntGen(t *testing.T) {
	for i := 0; i < 10; i++ {
		//fmt.Printf(fmt.Sprintf("%d\n", Fake.Int64Between(int64(0), int64(9223372036854775807))))
		fmt.Printf(fmt.Sprintf("%f\n", Fake.Float32(4, 10, 2)))
	}
}
