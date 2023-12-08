package schema2json

import (
	"reflect"
	"testing"
)

func TestGenerateArray(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   []interface{}
	}
	tests := []test{
		{
			name: "default",
			schema: &Schema{
				Type: stringPtr("list"),
				Items: &Schema{
					Type: stringPtr("string"),
				},
				Default: []interface{}{"bar"},
			},
			want: []interface{}{"bar"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateArray(tc.schema)
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
