package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Payload struct {
	// Environment variables to set for the deploy process.
	Env map[string]string
}

// ReadPayload takes a path to the JSON payload file and parses it
// into a Payload struct.
func ReadPayload(path string) *Payload {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(raw, &p)
	if err != nil {
		panic(err)
	}

	return &p
}

func main() {
	var (
		// Flags
		payload = flag.String("payload", "", "Json payload string.")
		_       = flag.String("d", "", "Unkown d")
		_       = flag.String("e", "", "Unkown e")
		_       = flag.String("id", "", "Unkown id")
	)
	flag.Parse()

	p := ReadPayload(*payload)

	fmt.Println(p)
}
