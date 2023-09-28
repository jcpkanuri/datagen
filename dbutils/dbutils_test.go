package dbutils

import (
	"fmt"
	"testing"
)

func TestFetchDistinctValues(t *testing.T) {
	testObj, _ := FetchDistinctValues("mysinglestore", "store", "id")
	fmt.Println("result: ", testObj)
}
