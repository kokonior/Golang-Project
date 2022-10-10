package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func TestJsonTag(t *testing.T) {
	laptop := Product{
		Id:       "P-001",
		Name:     "Macbook Pro 15",
		ImageUrl: "https://images/image.png",
	}

	// marhsall
	jsonBytes, err := json.Marshal(laptop)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

	// unmarshall
	jsonExample := `{"id":"P-001","name":"Macbook Pro 15","imageUrl":"https://images/image.png"}`
	bytesJson := []byte(jsonExample)

	product := new(Product)

	err = json.Unmarshal(bytesJson, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(*product)

}
