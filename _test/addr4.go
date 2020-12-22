package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const jsonData = `[
  "foo",
  "bar"
]`

const jsonData2 = `[
  {"foo": "foo"},
  {"bar": "bar"}
]`

const jsonData3 = `{
  "foo": "foo",
  "bar": "bar"
}`

func fromSlice() {
	var a []interface{}
	var c, d interface{}
	c = 2
	d = 3
	a = []interface{}{c, d}

	if err := json.Unmarshal([]byte(jsonData), &a); err != nil {
		log.Fatalln(err)
	}

	for k, v := range a {
		fmt.Println(k, ":", v)
	}
}

func fromEmpty() {
	var a interface{}
	var c, d interface{}
	c = 2
	d = 3
	a = []interface{}{c, d}

	if err := json.Unmarshal([]byte(jsonData), &a); err != nil {
		log.Fatalln(err)
	}

	b := a.([]interface{})

	for k, v := range b {
		fmt.Println(k, ":", v)
	}
}

func sliceOfObjects() {
	var a interface{}

	if err := json.Unmarshal([]byte(jsonData2), &a); err != nil {
		log.Fatalln(err)
	}

	b := a.([]interface{})

	for k, v := range b {
		fmt.Println(k, ":", v)
	}
}

func intoMap() {
	var a interface{}

	if err := json.Unmarshal([]byte(jsonData3), &a); err != nil {
		log.Fatalln(err)
	}

	b := a.(map[string]interface{})

	// TODO(mpl): make it sorted or something.
	for k, v := range b {
		vv := v.(string)
		fmt.Println(k, ":", vv)
	}

	for k, v := range b {
		fmt.Println(k, ":", v)
	}
}

func main() {
	fromSlice()
	fromEmpty()
	sliceOfObjects()
	intoMap()
}

// Ouput:
// 0 : foo
// 1 : bar
// 0 : foo
// 1 : bar
// 0 : map[foo:foo]
// 1 : map[bar:bar]
// foo : foo
// bar : bar
// foo : foo
// bar : bar
