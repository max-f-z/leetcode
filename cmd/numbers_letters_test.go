package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type person struct {
	Name string `json:"name"`
}

func TestNumbersLetters(t *testing.T) {
	numsLetters()

	js := `{"name":"alex"}`
	var p person
	err := json.Unmarshal([]byte(js), &p)

	if err != nil {
		t.Log(err)
	}

	fmt.Println(p)
}
