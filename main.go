package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	if len(os.Args) < 2 {
		panic("no input given")
	}
	input := os.Args[1]

	bytes, err := ioutil.ReadFile(input)

	if err != nil {
		panic(err)
	}

	var asd interface{}

	err = json.Unmarshal(bytes, &asd)
	if err != nil {
		panic(err)
	}

	out, err := yaml.Marshal(asd)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(out)
}
