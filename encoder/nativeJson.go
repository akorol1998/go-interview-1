package encoder

import (
	"bytes"
	"errors"
	"fmt"
)

type BufferWrite func(buf *bytes.Buffer, params ...any)

type JsonStruct struct {
	Type   int
	Result []string
}

func escape(str string) string {
	var buf bytes.Buffer
	for i := range str {
		switch str[i] {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		case '\b':
			buf.WriteString(`\b`)
		case '\f':
			buf.WriteString(`\f`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func iter(buf *bytes.Buffer, jsStruct *JsonStruct, bufW BufferWrite) {
	for idx, str := range jsStruct.Result {
		bufW(buf, idx, str)
		if idx < len(jsStruct.Result)-1 {
			buf.WriteByte(',')
		}
	}
}

func bufFirstWriter(buf *bytes.Buffer, params ...any) {
	buf.WriteByte('"')
	buf.WriteString(escape(params[1].(string)))
	buf.WriteByte('"')
}

func bufSecondWriter(buf *bytes.Buffer, params ...any) {
	buf.WriteByte('"')
	buf.WriteByte(byte(params[0].(int) + 48)) // implicit conversion
	buf.WriteString(`":`)
	bufFirstWriter(buf, params...)
}

func JsonEncode(v *JsonStruct) ([]byte, error) {
	var bufW BufferWrite
	var suf byte

	if v == nil {
		return nil, errors.New("Can not encode nil value")
	}
	buf := bytes.NewBuffer(
		[]byte(fmt.Sprintf(`{"type":%d, "result":`, v.Type)),
	)
	if v.Type == 1 {
		buf.WriteByte('[')
		suf = ']'
		bufW = bufFirstWriter

	} else {
		buf.WriteByte('{')
		suf = '}'
		bufW = bufSecondWriter
	}

	iter(buf, v, bufW)
	buf.WriteByte(suf)
	buf.WriteByte('}')
	return []byte(buf.Bytes()), nil
}
