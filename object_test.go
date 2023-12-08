package schema2json

import (
	"reflect"
	"testing"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func TestGenerateObject(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   map[string]interface{}
	}
	tests := []test{
		{
			name: "basic",
			schema: &Schema{
				Type: stringPtr("object"),
				Properties: orderedmap.New[string, *Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *Schema]{
						Key: "name",
						Value: &Schema{
							Type:  stringPtr("string"),
							Const: "bob",
						},
					},
				)),
			},
			want: map[string]interface{}{
				"name": "bob",
			},
		},
		{
			name: "dependencies",
			schema: &Schema{
				Type: stringPtr("object"),
				Properties: orderedmap.New[string, *Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *Schema]{
						Key: "enable",
						Value: &Schema{
							Type: stringPtr("boolean"),
						},
					},
				)),
				Dependencies: map[string]*Schema{
					"enable": {
						OneOf: []*Schema{
							{
								Properties: orderedmap.New[string, *Schema](orderedmap.WithInitialData(
									orderedmap.Pair[string, *Schema]{
										Key: "enable",
										Value: &Schema{
											Const: true,
										},
									},
									orderedmap.Pair[string, *Schema]{
										Key: "extra",
										Value: &Schema{
											Type:  stringPtr("string"),
											Const: "foo",
										},
									},
								)),
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"enable": true,
				"extra":  "foo",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateObject(tc.schema)
			_ = got
			if err != nil {
				t.Fatalf("%s, unexpected error", err.Error())
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got: %v want: %v", got, tc.want)
			}
		})
	}
}
