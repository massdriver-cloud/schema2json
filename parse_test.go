package schema2json_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/massdriver-cloud/schema2json"
)

func stringPtr(value string) *string {
	return &value
}
func floatPtr(value float64) *float64 {
	return &value
}

func TestParse(t *testing.T) {
	type test struct {
		name     string
		filePath string
		want     *schema2json.Schema
	}
	tests := []test{
		{
			name:     "ParseFile",
			filePath: "./testdata/testschema.json",
			want: &schema2json.Schema{
				ID:      stringPtr("https://example.com/person.schema.json"),
				Version: stringPtr("http://json-schema.org/draft-07/schema#"),
				Title:   stringPtr("Person"),
				Properties: map[string]*schema2json.Schema{
					"name": {
						Type:        stringPtr("string"),
						Description: stringPtr("a name"),
						Enum:        []interface{}{"Bob", "Dan"},
					},
					"age": {
						Description: stringPtr("an integer with min/max and multipleOf"),
						Type:        stringPtr("integer"),
						Minimum:     floatPtr(0),
						Maximum:     floatPtr(10),
						MultipleOf:  floatPtr(3),
					},
					"float": {
						Description: stringPtr("A floating point value"),
						Type:        stringPtr("number"),
						Minimum:     floatPtr(-2341.5432),
						Maximum:     floatPtr(5423.1512345),
					},
					"hmmph": {
						Type:  stringPtr("integer"),
						Const: float64(20),
					},
					"object": {
						Title: stringPtr("test object"),
						Type:  stringPtr("object"),
						Properties: map[string]*schema2json.Schema{
							"nested": {
								Type: stringPtr("string"),
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := os.ReadFile(tc.filePath)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			input := bytes.NewBuffer(raw)
			got, err := schema2json.Parse(input)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
