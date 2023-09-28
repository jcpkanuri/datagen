package config

import (
	"datagen/dbutils"
	"datagen/types"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Enrich(conname string, t types.Table) types.Table {

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
