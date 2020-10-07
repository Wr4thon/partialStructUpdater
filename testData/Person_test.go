package testdata_test

import (
	"testing"

	testdata "github.com/Wr4thon/partialStructUpdater/testData"
)

const (
	baz = "Baz"
)

func Test_PersonUpdate(t *testing.T) {
	person := testdata.Person{
		Firstname:    "Foo",
		Lastname:     "Bar",
		FavoriteFood: testdata.Burger,
	}

	person.Update(testdata.PersonUpdate{
		Person: testdata.Person{
			Firstname: baz,
		},
		UpdateMask: testdata.PersonUpdateMaskFirstname,
	})

	if person.Firstname != baz {
		t.Logf("Firstname was not equal. expected: %v, actual: %v\n", baz, person.Firstname)
		t.FailNow()
	}
}
