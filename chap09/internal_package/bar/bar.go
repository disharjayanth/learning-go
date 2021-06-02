package bar

import (
	"fmt"

	"github.com/disharjayanth/learningGo/chap09/internal_package/foo/internal"
)

func PrintDoubler() {
	fmt.Println(internal.Doubler(2))
}
