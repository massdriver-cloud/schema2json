package schema2json

import (
	"reflect"
	"testing"
)

func TestGenerateBoolean(t *testing.T) {
	type test struct {
		name   string
		schema *Schema
		want   bool
	}
	tests := []test{
		{
			name: "default",
			schema: &Schema{
				Type:    stringPtr("boolean"),
				Default: true,
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateBoolean(tc.schema)
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
