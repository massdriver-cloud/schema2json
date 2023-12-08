package schema2json

import (
	"reflect"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   float64
	}
	tests := []test{
		{
			name: "default",
			schema: &Schema{
				Type:    stringPtr("number"),
				Default: 234.5,
				Enum:    []interface{}{1.1},
			},
			want: 234.5,
		},
		{
			name: "enum",
			schema: &Schema{
				Type: stringPtr("number"),
				Enum: []interface{}{1.1},
			},
			want: 1.1,
		},
		{
			name: "min/max",
			schema: &Schema{
				Type:    stringPtr("number"),
				Minimum: floatPtr(2),
				Maximum: floatPtr(2),
			},
			want: 2,
		},
		{
			name: "multiple",
			schema: &Schema{
				Type:       stringPtr("number"),
				Minimum:    floatPtr(2),
				Maximum:    floatPtr(5),
				MultipleOf: floatPtr(3),
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateNumber(tc.schema)
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
