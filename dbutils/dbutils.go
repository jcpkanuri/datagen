package dbutils

import (
	"context"
	"database/sql"
	"datagen/appconfig"
	"datagen/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const mysql_distinct_agg_query = "select group_concat(distinct %s ) from %s "

const pg_distinct_agg_query = "select string_agg(distinct %s, ',') from %s "

func GetDBConnection(conname string) (*sql.DB, error) {

	mydbcon := getDBConnDetails(conname)
	db, err := sql.Open(mydbcon.DBType, getConnectionString(mydbcon))

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func FetchDistinctValues(conname string, tableName string, columnName string) ([]string, error) {
	mydbcon := getDBConnDetails(conname)
	dbConn, _ := GetDBConnection(conname)

	defer dbConn.Close()
	var query string

	if mydbcon.DBType == "mysql" {
		query = fmt.Sprintf(mysql_distinct_agg_query, columnName, tableName)
	} else if mydbcon.DBType == "postgres" {
		query = fmt.Sprintf(pg_distinct_agg_query, columnName, tableName)
	} else {
		log.Fatal("Unkown database type")
	}

	//s := fmt.Sprintf("select group_concat(distinct %s ) from %s ", columnName, tableName)
	var result sql.NullString
	row := dbConn.QueryRow(query)
	err := row.Scan(&result)
	if err != nil {

		switch err {
		case sql.ErrNoRows:
			result = sql.NullString{}
		default:
			log.Fatal("Failed to read data", "query", query, "error", err)
		}
	}

	a := strings.Split(result.String, ",")
	return a, nil

}

func ExecuteInTransaction(conname string, queries []string, batchSize int) {
	chunks := chunkSlice(queries, batchSize)
	dbConn, _ := GetDBConnection(conname)

	defer dbConn.Close()
	ctx := context.Background()

	for index, batch := range chunks {
		tx, err := dbConn.BeginTx(ctx, nil)
		if err != nil {
			log.Fatal("Failed to initiate transaction")
		}

		for _, line := range batch {
			_, err := tx.ExecContext(ctx, line)

			if err != nil {
				log.Errorf("Encountered error while line: %d  err: %v\n", index, err)
				log.Error("-------------------------------------------")
				log.Error(line)
				log.Error("-------------------------------------------")
				tx.Rollback()
				return
			}
		}
		err = tx.Commit()

		if err != nil {
			log.Fatal(err)
		}

		if index > 0 && index%batchSize == 0 {
			log.Infof("executed %d statements", index)
		}

		log.Infof("%d Batch of insert statments of size %d complete", index, batchSize)
	}

	log.Infof("execution of %d insert statements completed.", len(queries))
}

func getConnectionString(m types.DbConn) string {
	switch m.DBType {
	case "mysql":
		connstr := m.User + ":" + m.Pass + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DbName
		return connstr
	case "postgres":
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", m.Host, m.Port, m.User, m.Pass, m.DbName)
		return psqlconn

	default:
		return ""
	}
}

func getDBConnDetails(conname string) types.DbConn {
	aconfig := appconfig.GetConf()
	var mydbcon types.DbConn

	for _, a := range aconfig.Conns {
		if a.ConName == conname {
			mydbcon = a
		}
	}

	return mydbcon
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
