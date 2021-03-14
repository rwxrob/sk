package main

import (
	"encoding/json"
	"fmt"
)

func ConvertToJSON(thing interface{}) string {
	byt, err := json.MarshalIndent(thing, "", "  ")
	if err != nil {
		return fmt.Sprintf("{\"ERROR\": \"%v\"}", err)
	}
	return string(byt)
}
