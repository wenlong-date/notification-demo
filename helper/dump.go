package helper

import (
	"encoding/json"
	"fmt"
)

func Dump(data ...any) {
	for _, v := range data {
		fmt.Printf("%+v\n", v)
	}
}

func DumpJson(data ...any) {
	for _, v := range data {
		j, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			fmt.Printf("json marshal err: %+v\n", v)
			continue
		}

		fmt.Println(string(j))
	}
}
