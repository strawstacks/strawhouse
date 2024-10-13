package common

import (
	"encoding/json"
	"fmt"
	"github.com/bsthun/gut"
)

func Handle(data any, er error) {
	if er != nil {
		gut.Error("command failed", er, true)
		return
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		gut.Fatal("Unable to marshal error", err)
	}
	fmt.Println(string(bytes))
}
