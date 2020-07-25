package helper

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonPretty(values interface{}) {
	prettyJSON, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
