package generators

import (
	"fmt"
	"testing"

	"slices"
)

func TestFirstName(t *testing.T) {

	got := Fake.Person().FirstName()
	a := Fake.Json()
	l := a.String()
	fmt.Println(l)

	if len(got) == 0 {
		t.Errorf("got %q is empty", got)
	}
}

// Map datatype to specific faker func to generate data

func TestFakeRandomNess(t *testing.T) {
	var container []string
	i := 1
	for {
		item := Fake.Person().FirstName()

		if !slices.Contains(container, item) {
			container = append(container, item)
			//fmt.Print("entry added\n")
			i -= 1
		} else {
			//fmt.Printf("found match at %d among %d items\n", slices.Index(container, item), len(container))
			i += 1
		}

		if i > 1000 {

			break
		}

	}

	fmt.Printf("out of loop %d\n", len(container))

}
