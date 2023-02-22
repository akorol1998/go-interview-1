package main

import (
	"fmt"
)

type jsonStruct struct {
	Type   int
	Result []string
}

func (s *jsonStruct) jsonEncode() []byte {
	var suf string

	jsonString := fmt.Sprintf(`{"type":%d, "result":[`, s.Type)
	suf = `]`
	if s.Type == 2 {
		jsonString = jsonString[:len(jsonString)-1] + "{"
		suf = `}`
	}
	for idx, str := range s.Result {
		if s.Type == 2 {
			jsonString += fmt.Sprintf(`"%d":"%s",`, idx, str)
		} else {
			jsonString += fmt.Sprintf(`"%s",`, str)
		}
	}

	jsonString = jsonString[:len(jsonString)-1]
	jsonString += fmt.Sprintf(`%s]`, suf)
	return []byte(jsonString)
}

func main() {
	num := 10
	s1 := make([]string, num)
	for i := range s1 {
		s1[i] = fmt.Sprintf("res-%d", i)
	}
	s2 := make([]string, num)
	copy(s2, s1)
	js1 := jsonStruct{1, s1}
	js2 := jsonStruct{2, s2}

	fmt.Printf("json encoded type:1: %s\n", js1.jsonEncode())
	fmt.Printf("json encoded type:2: %s\n", js2.jsonEncode())
}
