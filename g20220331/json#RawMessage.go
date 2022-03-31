package main

// https://stackoverflow.com/questions/19145202/marshal-of-json-rawmessage
// json.RawMessage is a []byte, but it can be used to store any type of data.
import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name string
	Id   int
	Json *json.RawMessage
}
type Data2 struct {
	Name string
	Id   int
}

func main() {

	tmp := Data2{"World", 2}

	b, err := json.Marshal(tmp)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}
	fmt.Printf("b %s", string(b))

	test := Data{"Hello", 1, (*json.RawMessage)(&b)}
	b2, err := json.Marshal(test)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}

	fmt.Printf("b2 %s", string(b2))

	var d Data
	err = json.Unmarshal(b2, &d)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}
	fmt.Printf("d.Json %s", string(*d.Json))

	var tmp2 Data2
	err = json.Unmarshal(*d.Json, &tmp2)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}

	fmt.Printf("Data2 %+v", tmp2)
}
