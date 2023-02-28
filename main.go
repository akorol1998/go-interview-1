package main

import (
	"encoding/json"
	"fmt"

	"github.com/akorol1998/go-interview-1/encoder"
)

func main() {
	num := 10
	s1 := make([]string, num)
	for i := range s1 {
		s1[i] = fmt.Sprintf(`res-%s-%v`, "\t--\n-\\-\b-\f-\r", i)
	}
	s2 := make([]string, num)
	copy(s2, s1)
	js1 := encoder.JsonStruct{1, s1}
	js2 := encoder.JsonStruct{2, s2}

	encJs1, _ := encoder.JsonEncode(&js1)
	encJs2, _ := encoder.JsonEncode(&js2)

	fmt.Printf("json encoded type:1\nencoded:%s is_valid:%v\n", encJs1, json.Valid(encJs1))
	fmt.Printf("json encoded type:2\nencoded:%s is_valid:%v\n", encJs2, json.Valid(encJs2))
}
