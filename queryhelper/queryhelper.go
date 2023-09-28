package queryhelper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"

	"datagen/types"
)

func GetInsertStatementPrefix(t types.Table) (string, error) {

	var contentBuilder strings.Builder

	contentBuilder.WriteString(fmt.Sprintf("INSERT into %s (", t.Name))

	for _, v := range t.Columns {

		if strings.Contains(v.Extra, "increment") {
			continue
		}

		contentBuilder.WriteString(v.Name + ",")
	}

	result := strings.TrimSuffix(contentBuilder.String(), ",")

	result += ") values "

	return result, nil
}

func GetRow(t types.Table, rowIndex int) (types.Row, error) {

	colsize := len(t.Columns) - len(t.AutoIncCols)

	var result = make([]string, colsize)
	valueIndex := 0

	for _, c := range t.Columns {

		if strings.Contains(c.Extra, "increment") {
			continue
		} else if c.Sequence {
			nextSeqVal, err := t.ColSequences.IncrementAndGet(c.Name, c.SeqConfig)
			if err != nil {
				log.Fatal("Sequence failed.", err)
			}
			result[valueIndex] = strconv.Itoa(nextSeqVal)
		} else {
			result[valueIndex] = handleValueGeneration(c, t)
		}
		valueIndex += 1

	}

	return types.Row{Pos: uint(rowIndex), ColValues: result}, nil
}

func GetRows(t types.Table, rowcount int) ([]types.Row, error) {
	var result = make([]types.Row, rowcount)

	for i := 0; i < rowcount; i++ {
		rowItem, err := GetRow(t, i)

		if err != nil {
			log.Fatal("Error: ", err)
		} else {
			result[i] = rowItem
		}
	}

	return result, nil
}

func GenerateFullInsertStatement(t types.Table, rowcount int) (string, error) {

	var sb strings.Builder

	result, err := GetInsertStatementPrefix(t)

	if err != nil {
		log.Fatal("Error --> ", err)
	}

	sb.WriteString(result)

	values, err := GetRows(t, rowcount)

	if err != nil {
		log.Fatal("Error --> ", err)
	}

	for _, v := range values {
		sb.WriteString(v.ValuesString())
		sb.WriteString(",")
	}

	return strings.TrimSuffix(sb.String(), ","), nil

}

// unqMap map[string][]string, colCardMap map[string][]string
func handleValueGeneration(c types.Column, t types.Table) string {

	var item string
	if c.Key == "UNI" {
		item = c.GenerateNewValue()
		limit := 999
		for limit > 0 && slices.Contains(t.UniqCols[c.Name], item) {
			item = c.GenerateNewValue()
			limit -= 1
		}
		t.UniqCols[c.Name] = append(t.UniqCols[c.Name], item)
	} else if c.Cardinality > 0 {
		if len(t.CardinalityCols[c.Name]) < c.Cardinality {
			item = c.GenerateNewValue()
			t.CardinalityCols[c.Name] = append(t.CardinalityCols[c.Name], item)
		} else {
			item = types.RandomStringElement(t.CardinalityCols[c.Name])
		}
	} else {
		item = c.GenerateNewValue()
	}

	return item
}
