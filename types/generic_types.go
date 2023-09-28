package types

import (
	"github.com/jaswdr/faker"
)

var Fake = faker.New()

type DbConn struct {
	ConName string `json:"conname"`
	DBType  string `json:"dbtype"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	DbName  string `json:"dbname"`
}

type Pair struct {
	First, Second int
}

func TruncateText(s string, max int) string {

	if len(s) > max {
		return s[:max]
	} else {
		return s
	}

}

func RandomStringElement(s []string) string {
	i := Fake.IntBetween(0, len(s)-1)
	return s[i]
}
