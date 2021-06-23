package comparefunc_test

import (
	"testing"
	"time"

	comparefunc "github.com/disharjayanth/learningGo/chap13/compareFunc"
	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := comparefunc.Person{
		Name:      "Dennis",
		Age:       37,
		DateAdded: time.Now(),
	}

	result := comparefunc.CreatePerson("Dennis", 37)

	comparer := cmp.Comparer(func(x, y comparefunc.Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
