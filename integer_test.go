package schema2json

import (
	"reflect"
	"testing"
)

func TestGenerateInteger(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   int
	}
	tests := []test{
		{
			name: "default",
			schema: &Schema{
				Type:    stringPtr("integer"),
				Default: 234.,
				Enum:    []interface{}{9.},
			},
			want: 234,
		},
		{
			name: "enum",
			schema: &Schema{
				Type: stringPtr("integer"),
				Enum: []interface{}{8.},
			},
			want: 8,
		},
		{
			name: "min/max",
			schema: &Schema{
				Type:    stringPtr("integer"),
				Minimum: floatPtr(2),
				Maximum: floatPtr(2),
			},
			want: 2,
		},
		{
			name: "multiple",
			schema: &Schema{
				Type:       stringPtr("integer"),
				Minimum:    floatPtr(2),
				Maximum:    floatPtr(5),
				MultipleOf: floatPtr(3),
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateInteger(tc.schema)
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
