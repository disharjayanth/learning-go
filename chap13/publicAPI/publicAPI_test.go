package publicapi_test

import (
	"testing"

	publicapi "github.com/disharjayanth/learningGo/chap13/publicAPI"
)

func TestAddNumbers(t *testing.T) {
	result := publicapi.AddNumbers(2, 3)
	if result != 5 {
		t.Errorf("Got %v Expected %v", result, 5)
	}
}
