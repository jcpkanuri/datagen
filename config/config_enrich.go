package config

import (
	"bytes"
	"datagen/dbutils"
	"datagen/types"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Enrich(conname string, t types.Table) types.Table {
	decimaltypes := []string{"float", "double", "decimal"}
	numericTypes := []string{"tinyint", "smallint", "mediumint", "int", "bigint"}

	t.UniqCols = make(map[string][]string)
	t.CardinalityCols = make(map[string][]string)
	t.ColSequences = make(map[string]int)

	for i, c := range t.Columns {
		//build datatype info map
		datatypeInfo, _ := ParseDataType(c.DataType)

		t.Columns[i].DataTypeName = datatypeInfo[0]
		if len(datatypeInfo) == 3 {
			t.Columns[i].DataTypeLimits = types.Pair{First: GetInt(datatypeInfo[1]), Second: GetInt(datatypeInfo[2])}
		} else if len(datatypeInfo) == 2 {
			t.Columns[i].DataTypeLimits = types.Pair{First: GetInt(datatypeInfo[1]), Second: -999}
		}

		if slices.Contains(numericTypes, datatypeInfo[0]) {
			t.Columns[i].NumericGenPattern = GenerateNumPattern(t.Columns[i].DataTypeLimits)
		} else if slices.Contains(decimaltypes, datatypeInfo[0]) {
			t.Columns[i].NumericGenPattern = GenerateDecimalPattern(t.Columns[i].DataTypeLimits)
		}

		//build uniq col map
		if c.Key == "UNI" {
			t.UniqCols[c.Name] = []string{}
		}

		if strings.Contains(c.Extra, "increment") {
			t.AutoIncCols = append(t.AutoIncCols, c.Name)
		}

		//build items for oneof generator
		item := c.Generator
		//Build oneof values
		if (item.Name == "oneof") && len(item.Items) == 0 {

			foreignKey := item.ForeignKey
			myItems, err := dbutils.FetchDistinctValues(foreignKey.Conname, foreignKey.Table, foreignKey.Column)

			if err == nil && len(myItems) > 0 {
				log.Debugf("Loading %d items from reference table %s and column %s", len(myItems), foreignKey.Table, foreignKey.Column)
				item.Items = myItems
				t.Columns[i].Generator = item
			}
		}

		if c.Key != "UNI" && c.Cardinality > 0 {
			t.CardinalityCols[c.Name] = []string{}
		}

		if c.Sequence {
			t.ColSequences[c.Name] = c.SeqConfig.Start
		}

		log.Debug("Column: %s , Definition: %+v\n", t.Columns[i].Name, t.Columns[i])
	}
	return t
}

func ParseDataType(input string) ([]string, error) {
	return strings.FieldsFunc(input, func(r rune) bool {
		switch r {
		case '(', ')', ',':
			return true
		}
		return false
	}), nil
}

func GetInt(input string) int {
	len, _ := strconv.Atoi(input)
	return len
}

func GenerateDecimalPattern(s types.Pair) string {
	var b bytes.Buffer

	for i := 1; i <= s.First; i++ {
		b.WriteString("#")
		if (s.Second != -999) && i == (s.First-s.Second) {
			b.WriteString(".")
		}
	}
	return b.String()
}

func GenerateNumPattern(s types.Pair) string {
	var b bytes.Buffer

	for i := 1; i <= s.First; i++ {
		b.WriteString("#")
	}
	return b.String()
}
