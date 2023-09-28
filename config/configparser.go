package config

import (
	"bytes"
	"encoding/json"

	"fmt"
	"os"

	"github.com/charmbracelet/log"

	"datagen/dbutils"
	"datagen/types"

	"gopkg.in/yaml.v3"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func ProbeTableMetadata(conname string, tableName string) (types.Table, error) {

	db, err := dbutils.GetDBConnection(conname)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//res, err := db.Query(fmt.Sprintf(`desc %s;`, "fact_sales"))

	var myTable types.Table

	myTable.Name = tableName

	res, err := db.Query(fmt.Sprintf(`select column_name as FieldName,
	column_type as FieldType,
	is_nullable as FieldNull,
	column_key as FieldKey,
	COALESCE(column_default,'') as FieldDefault,
	COALESCE(extra, '') as EXTRA 
	from information_schema.columns where table_name =  '%s';`, tableName))

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	mycols := []types.Column{}

	for res.Next() {

		var col types.Column
		err := res.Scan(&col.Name, &col.DataType, &col.Nullable, &col.Key,
			&col.DefaultVal, &col.Extra)

		if err != nil {
			log.Fatal(err)
		}

		mycols = append(mycols, col)
	}
	myTable.Columns = mycols

	return myTable, err
}

func ParseJsonFile(inFile string) (types.Table, error) {
	log.Infof("about to read File: %s", inFile)

	file, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Unable to open provided file: %s, %v ", inFile, err)
	}
	defer file.Close()

	var t types.Table

	log.Debug("file exists, about to parse json")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&t)

	if err != nil {
		log.Fatal("Failed to parse json", err)
	}

	log.Info("parse completed")
	return t, nil
}

func ParseYamlFile(inFile string) (types.Table, error) {
	log.Infof("about to read File: %s", inFile)
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Unable to open provided file: %s", inFile)
	}
	defer file.Close()

	var t types.Table
	log.Debug("file exists, about to parse yaml")
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&t)

	if err != nil {
		log.Fatalf("Failed to parse yaml. %v", err)
	}
	log.Info("parse completed")
	return t, nil
}
