package generators

import (
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/jaswdr/faker"
)

var Fake = faker.New()

type Generator struct {
	Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
	Min        int            `json:"min,omitempty" yaml:"min,omitempty"`
	Max        int            `json:"max,omitempty" yaml:"max,omitempty"`
	Len        int            `json:"length,omitempty" yaml:"length,omitempty"`
	Expression string         `json:"exp,omitempty" yaml:"exp,omitempty"`
	Items      []string       `json:"items,omitempty" yaml:"items,omitempty"`
	ForeignKey ForeignMapping `json:"foreignKey,omitempty" yaml:"foreignKey,omitempty"`
}

type ForeignMapping struct {
	Table   string `json:"table,omitempty" yaml:"table,omitempty"`
	Column  string `json:"column,omitempty" yaml:"column,omitempty"`
	Conname string `json:"conname,omitempty" yaml:"conname,omitempty"`
}

func (g Generator) Generate() string {

	var retVal string
	switch g.Name {
	case "string":
		retVal = Fake.RandomStringWithLength(g.Len)
	case "firstname":
		retVal = Fake.Person().FirstName()
	case "lastname":
		retVal = Fake.Person().LastName()
	case "fullname":
		retVal = Fake.Person().Name()
	case "lexify":
		retVal = Fake.Lexify(g.Expression)
	case "numerify":
		retVal = Fake.Numerify(g.Expression)
	case "bothify":
		retVal = Fake.Bothify(g.Expression)
	case "oneof":
		if len(g.Items) > 0 {
			retVal = Fake.RandomStringElement(g.Items)
		} else {
			fmt.Println("length of items is empty")
			retVal = "" //todo check empty string implications
		}

	default:
		log.Fatalf("Generator: %s is not recognized", g.Name)
	}

	return retVal
}
