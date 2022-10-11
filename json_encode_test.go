package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJson(t *testing.T) {
	type Website struct {
		WebName string
		Link    string
	}

	type Person struct {
		Name    string
		Age     int
		Married bool
		Hobby   []string
		Website Website
	}

	neo := Person{
		Name:    "Neo Jarmawijaya",
		Age:     20,
		Married: false,
		Hobby: []string{
			"Coding",
			"Gaming",
			"Reading",
		},
		Website: Website{
			WebName: "neo blog",
			Link:    "https://neoj.com",
		},
	}

	// marshall
	bytes, err := json.Marshal(neo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	// unmarshall
	var obj Person
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)
}
