package schema2json_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/massdriver-cloud/schema2json"
	"github.com/stretchr/testify/require"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func stringPtr(value string) *string {
	return &value
}
func floatPtr(value float64) *float64 {
	return &value
}

func TestMarshal(t *testing.T) {
	type testData struct {
		name   string
		schema schema2json.Schema
	}
	tests := []testData{
		{
			name: "addprop",
			schema: schema2json.Schema{
				Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropFalse",
						Value: &schema2json.Schema{
							Type:                 stringPtr("object"),
							AdditionalProperties: false,
							Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
								orderedmap.Pair[string, *schema2json.Schema]{
									Key: "foo",
									Value: &schema2json.Schema{
										Type: stringPtr("string"),
									},
								},
							)),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropTrue",
						Value: &schema2json.Schema{
							Type:                 stringPtr("object"),
							AdditionalProperties: true,
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropSchema",
						Value: &schema2json.Schema{
							Type: stringPtr("object"),
							AdditionalProperties: &schema2json.Schema{
								Type: stringPtr("object"),
								Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
									orderedmap.Pair[string, *schema2json.Schema]{
										Key: "bar",
										Value: &schema2json.Schema{
											Type: stringPtr("string"),
										},
									},
								)),
							},
						},
					},
				)),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want, err := os.ReadFile(filepath.Join("testdata", tc.name+".json"))
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			got, err := json.MarshalIndent(tc.schema, "", "    ")
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			require.JSONEq(t, string(got), string(want))
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type testData struct {
		name string
		want schema2json.Schema
	}
	tests := []testData{
		{
			name: "addprop",
			want: schema2json.Schema{
				Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropFalse",
						Value: &schema2json.Schema{
							Type:                 stringPtr("object"),
							AdditionalProperties: false,
							Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
								orderedmap.Pair[string, *schema2json.Schema]{
									Key: "foo",
									Value: &schema2json.Schema{
										Type: stringPtr("string"),
									},
								},
							)),
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropTrue",
						Value: &schema2json.Schema{
							Type:                 stringPtr("object"),
							AdditionalProperties: true,
						},
					},
					orderedmap.Pair[string, *schema2json.Schema]{
						Key: "addPropSchema",
						Value: &schema2json.Schema{
							Type: stringPtr("object"),
							AdditionalProperties: &schema2json.Schema{
								Type: stringPtr("object"),
								Properties: orderedmap.New[string, *schema2json.Schema](orderedmap.WithInitialData(
									orderedmap.Pair[string, *schema2json.Schema]{
										Key: "bar",
										Value: &schema2json.Schema{
											Type: stringPtr("string"),
										},
									},
								)),
							},
						},
					},
				)),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bytes, err := os.ReadFile(filepath.Join("testdata", tc.name+".json"))
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			got := schema2json.Schema{}
			err = json.Unmarshal(bytes, &got)
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got: %#v want %#v", got, tc.want)
			}
		})
	}
}
