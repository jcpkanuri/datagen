package types

import (
	"datagen/generators"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type Table struct {
	Name            string              `json:"tableName" yaml:"tableName"`
	Columns         []Column            `json:"columns" yaml:"columns"`
	Rows            []Row               `json:"rows" yaml:"rows"`
	UniqCols        map[string][]string `json:"-" yaml:"-"`
	CardinalityCols map[string][]string `json:"-" yaml:"-"`
	AutoIncCols     []string            `json:"-" yaml:"-"`
	ColSequences    SequenceMap         `json:"-" yaml:"-"`
}

type Column struct {
	Name              string               `json:"name" yaml:"name"`
	DataType          string               `json:"type" yaml:"type"`
	Nullable          string               `json:"nullsAllowed" yaml:"nullsAllowed"`
	Key               string               `json:"key" yaml:"key"`
	DefaultVal        string               `json:"default" yaml:"default"`
	Extra             string               `json:"extra" yaml:"extra"`
	Generator         generators.Generator `json:"generator,omitempty" yaml:"generator,omitempty"`
	Sequence          bool                 `json:"sequence" yaml:"sequence"`
	DataTypeName      string               `json:"-" yaml:"-"`
	DataTypeLimits    Pair                 `json:"-" yaml:"-"`
	Cardinality       int                  `json:"cardinality" yaml:"cardinality"`
	SeqConfig         SequenceConfig       `json:"seqConfig" yaml:"seqConfig"`
	NumericGenPattern string               `json:"-" yaml:"-"`
}

type Row struct {
	Pos       uint
	ColValues []string
}

func (r Row) ValuesString() string {
	return "(" + strings.Join(r.ColValues, ",") + ")"
}

func (c Column) GenerateNewValue() string {
	var retVal string
	log.Debugf("-----------\nName: %s\nDataType:%s\nLimits: %+v\n", c.Name, c.DataTypeName, c.DataTypeLimits)
	if c.Generator.Name == "" {

		switch c.DataTypeName {
		//integer types
		case "tinyint":
			retVal = TruncateText(strconv.Itoa(int(Fake.Int8Between(-128, 127))), c.DataTypeLimits.First)
		case "smallint":
			retVal = TruncateText(strconv.Itoa(int(Fake.Int16Between(-32768, 32767))), c.DataTypeLimits.First)
		case "mediumint":
			retVal = TruncateText(strconv.Itoa(int(Fake.Int32Between(-8388608, 8388607))), c.DataTypeLimits.First)
		case "int":
			retVal = TruncateText(strconv.Itoa(int(Fake.Int32Between(-2147483648, 2147483647))), c.DataTypeLimits.First)
		case "bigint":
			retVal = TruncateText(strconv.FormatInt(Fake.Int64Between(0, 9223372036854775807), 10), c.DataTypeLimits.First)
			//float
		case "float", "double":
			retVal = fmt.Sprintf("%f", generateDecimal32(c.NumericGenPattern))
			//decimal
		case "decimal":
			retVal = fmt.Sprintf("%f", generateDecimal64(c.NumericGenPattern))
		case "char", "binary":
			retVal = "'" + Fake.Letter() + "'"
		case "varchar", "varbinary":
			//min, _ := strconv.Atoi(datatype[1])
			retVal = fmt.Sprintf("%s%s%s", "'", Fake.RandomStringWithLength(c.DataTypeLimits.First), "'")
		case "longtext":
			retVal = fmt.Sprintf("%s%s%s", "'", Fake.Lorem().Sentence(15), "'")
			//Todo -- add date time types
		case "date":
			retVal = "'" + Fake.Time().Time(time.Now()).Format("2006-01-02") + "'"
		case "time":
			retVal = "'" + Fake.Time().Time(time.Now()).Format("15:04:05") + "'"
		case "timestamp":
			retVal = "'" + Fake.Time().Time(time.Now()).Format("2006-01-02 15:04:05") + "'"
		case "year":
			retVal = fmt.Sprintf("%d", Fake.Time().Year())

		case "enum", "set":
			retVal = TruncateText(fmt.Sprintf("%d", Fake.Int8()), 2)
		case "geographypoint", "geography":
			retVal = fmt.Sprintf("GEOGRAPHY_POINT(%f,%f)", Fake.Address().Latitude(), Fake.Address().Longitude())
		// case "geography":
		// 	// start := fmt.Sprintf("%d %d", Fake.Int8Between(0, 9), Fake.Int8Between(0, 9))
		// 	// retVal = fmt.Sprintf("'POLYGON((%s ,%d %d,%d %d, %s))'",
		// 	// 	start, Fake.Int8Between(0, 9), Fake.Int8Between(0, 9), Fake.Int8Between(0, 9), Fake.Int8Between(0, 9), start)
		// 	retVal = RandomGeoPolygon()
		case "json":
			jsonItem := Fake.Json()
			retVal = jsonItem.String()
		}
	} else {
		retVal = fmt.Sprintf("%s%s%s", "'", TruncateText(c.Generator.Generate(), c.DataTypeLimits.First), "'")

	}

	return retVal
}

func generateDecimal32(inputPattern string) float32 {
	value, _ := strconv.ParseFloat(Fake.Numerify(inputPattern), 32)
	return float32(value)
}

func generateDecimal64(inputPattern string) float64 {
	return float64(generateDecimal32(inputPattern))
}
