package schema2json_test

import (
	"reflect"
	"testing"

	"github.com/massdriver-cloud/schema2json"
)

func TestMergeSchemas(t *testing.T) {
	type test struct {
		name  string
		base  *schema2json.Schema
		merge *schema2json.Schema
		want  *schema2json.Schema
	}
	tests := []test{
		{
			name: "Base schema retained",
			base: &schema2json.Schema{
				Type: stringPtr("string"),
			},
			merge: &schema2json.Schema{},
			want: &schema2json.Schema{
				Type: stringPtr("string"),
			},
		},
		{
			name: "Merge schema set",
			base: &schema2json.Schema{},
			merge: &schema2json.Schema{
				Type: stringPtr("string"),
			},
			want: &schema2json.Schema{
				Type: stringPtr("string"),
			},
		},
		{
			name: "Merge wins in collision",
			base: &schema2json.Schema{
				Type: stringPtr("string"),
			},
			merge: &schema2json.Schema{
				Type: stringPtr("number"),
			},
			want: &schema2json.Schema{
				Type: stringPtr("number"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := schema2json.MergeSchemas(tc.base, tc.merge)
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
