package gocmptocomparetestres_test

import (
	"testing"

	gocmptocomparetestres "github.com/disharjayanth/learningGo/chap13/goCmpToCompareTestRes"
	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := gocmptocomparetestres.Person{
		Name: "Dennis",
		Age:  38,
	}

	result := gocmptocomparetestres.CreatePerson("Dennis", 38)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
