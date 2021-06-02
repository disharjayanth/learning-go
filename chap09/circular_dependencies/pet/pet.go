package pet

import "github.com/learning-go-book/circular_dependency_example/person"

var owners = map[string]person.Person{
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}
