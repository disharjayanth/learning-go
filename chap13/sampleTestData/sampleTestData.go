package sampletestdata

import (
	"io/ioutil"
	"unicode/utf8"
)

func countCharacters(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
