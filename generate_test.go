package schema2json_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/massdriver-cloud/schema2json"
	"github.com/xeipuuv/gojsonschema"
)

func TestGenerateJSON(t *testing.T) {
	type test struct {
		name       string
		schemaPath string
	}
	tests := []test{
		{
			name:       "Foo",
			schemaPath: "testdata/testschema.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			schema := new(schema2json.Schema)
			bytes, err := os.ReadFile(tc.schemaPath)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}
			err = json.Unmarshal(bytes, schema)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			got, err := schema2json.GenerateJSON(schema)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			schemaLoader := gojsonschema.NewReferenceLoader("file://" + tc.schemaPath)
			documentLoader := gojsonschema.NewGoLoader(got)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}
			if !result.Valid() {
				for _, violation := range result.Errors() {
					t.Fatalf("\t- %v\n", violation)
				}
			}
		})
	}
}
