package schema2json

import (
	"bytes"
	"encoding/json"
	"io"
)

func Parse(input io.Reader) (*Schema, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)

	schema := new(Schema)
	err := json.Unmarshal(buf.Bytes(), schema)
	return schema, err
}
