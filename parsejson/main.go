// Parsejson parses a JSON string and prints it out.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var files = `[
    {
        "path": "/etc",
        "is_dir": true
    },
    {
        "path": "/etc/hosts",
        "is_dir": false
    }
]`

func main() {
	r := strings.NewReader(files)

	var fs []struct {
		Path  string
		IsDir bool `json:"is_dir"`
	}
	if err := json.NewDecoder(r).Decode(&fs); err != nil {
		log.Fatal(err)
	}

	for _, f := range fs {
		fmt.Printf("%-10s %t\n", f.Path, f.IsDir)
	}
}