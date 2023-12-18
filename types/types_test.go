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
