package main

import "fmt"

func main() {
	var nilMap map[string]int // nil map
	fmt.Println("nilMap:", nilMap)
	if nilMap == nil {
		fmt.Println("nilMap is nil.")
	}

	totalWins := map[string]int{}
	fmt.Println("totalWins:", totalWins)
	if totalWins == nil {
		fmt.Println("totalWins is nil.")
	}

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	for k, v := range teams {
		fmt.Printf("%v\t%v\n", k, v)
	}

	ages := make(map[int][]string, 10)
	fmt.Println(ages, len(ages))

	ages[20] = []string{"John"}
	fmt.Println(ages, len(ages))

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins)
	totalWins["Kittens"]++

	for k, v := range totalWins {
		fmt.Println(k, ":", v)
	}

	var m map[string]int
	m = map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println("map m:", m)

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println("map m after deleting:", m)
}
