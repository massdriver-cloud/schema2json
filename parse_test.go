package schema2json_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/massdriver-cloud/schema2json"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

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
				Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "name",
						Value: &schema2json.Schema{
							Type:        stringPtr("string"),
							Description: stringPtr("a name"),
							Enum:        []interface{}{"Bob", "Dan"},
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "age",
						Value: &schema2json.Schema{
							Description: stringPtr("an integer with min/max and multipleOf"),
							Type:        stringPtr("integer"),
							Minimum:     floatPtr(0),
							Maximum:     floatPtr(10),
							MultipleOf:  floatPtr(3),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "float",
						Value: &schema2json.Schema{
							Description: stringPtr("A floating point value"),
							Type:        stringPtr("number"),
							Minimum:     floatPtr(-2341.5432),
							Maximum:     floatPtr(5423.1512345),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "hmmph",
						Value: &schema2json.Schema{
							Type:  stringPtr("integer"),
							Const: float64(20),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "object",
						Value: &schema2json.Schema{
							Title: stringPtr("test object"),
							Type:  stringPtr("object"),
							Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
								orderedmap.Pair[string, *schema2json.Schema]{
									Key: "nested",
									Value: &schema2json.Schema{
										Type: stringPtr("string"),
									},
								},
							)),
						},
					},
				)),
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

func TestParseMapStringInterface(t *testing.T) {
	type test struct {
		name  string
		input map[string]interface{}
		want  *schema2json.Schema
	}
	tests := []test{
		{
			name: "ParseMapStringInterface",
			input: map[string]interface{}{
				"$id":     "http://some-id.json",
				"$schema": "http://path.org/draft-07/schema",
				"title":   "example",
				"properties": map[string]interface{}{
					"age": map[string]interface{}{
						"type":        "integer",
						"description": "an integer field",
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "a string name",
					},
				},
			},
			want: &schema2json.Schema{
				ID:      stringPtr("http://some-id.json"),
				Version: stringPtr("http://path.org/draft-07/schema"),
				Title:   stringPtr("example"),
				Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "age",
						Value: &schema2json.Schema{
							Type:        stringPtr("integer"),
							Description: stringPtr("an integer field"),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "name",
						Value: &schema2json.Schema{
							Type:        stringPtr("string"),
							Description: stringPtr("a string name"),
						},
					},
				)),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := schema2json.ParseMapStringInterface(tc.input)
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
