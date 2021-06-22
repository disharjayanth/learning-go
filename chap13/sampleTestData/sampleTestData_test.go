package sampletestdata

import "testing"

func TestCountCharacters(t *testing.T) {
	total, err := countCharacters("testdata/sampletextFile.txt")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if total != 11 {
		t.Error("Expected 11 got", total)
	}

	_, err = countCharacters("testdata/nofile.txt")
	if err == nil {
		t.Error("Expected error!")
	}
}
