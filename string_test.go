package schema2json

import (
	"reflect"
	"testing"
)

func TestGenerateString(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   string
	}
	tests := []test{
		{
			name: "default",
			schema: &Schema{
				Type:    stringPtr("string"),
				Default: "default",
				Enum:    []interface{}{"other"},
			},
			want: "default",
		},
		{
			name: "enum",
			schema: &Schema{
				Type: stringPtr("string"),
				Enum: []interface{}{"enum"},
			},
			want: "enum",
		},
		{
			name: "pattern",
			schema: &Schema{
				Type:    stringPtr("string"),
				Pattern: stringPtr("^pattern$"),
			},
			want: "pattern",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateString(tc.schema)
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
