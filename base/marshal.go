package base

import (
	"fmt"

	"sigs.k8s.io/yaml"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type AUser struct {
	Name string
	Age  int
	test string
}

func Marshal() {
	body := struct {
		Name           string
		Id             int
		expectedStatus string
		expectedBody   string
	}{
		Name: "aaa",
		Id:   55,
	}
	/*body := AUser{
		Name: "aaa",
		Age:  55,
		test: "test",
	}*/

	b, err := yaml.Marshal(body)
	fmt.Printf("Get marshal result:%s, %s..\n", string(b), err)
}
