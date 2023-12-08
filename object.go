package schema2json

import (
	"encoding/json"
	"fmt"
	"maps"
)

func generateObject(property *Schema) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	var err error

	for prop := property.Properties.Oldest(); prop != nil; prop = prop.Next() {
		property.Name = prop.Key
		result[prop.Key], err = generateValue(prop.Value)
		if err != nil {
			return result, err
		}
	}

	if property.Dependencies != nil {
		for _, schema := range property.Dependencies {
			//dep := map[string]interface{}{}
			dep, err := generateValue(schema)
			s, _ := json.Marshal(result)
			fmt.Println(string(s))
			b, _ := json.Marshal(dep)
			fmt.Println(string(b))
			// var err error
			// property, err = MergeSchemas(property, schema)
			if err != nil {
				return result, err
			}
			// a, _ := json.Marshal(property)
			// fmt.Println(string(a))
			//applyDependency(name, property, schema)
			maps.Copy(result, dep.(map[string]interface{}))

		}
	}

	return result, nil
}

func applyDependency(propertyName string, current, dependencySchema *Schema) error {
	return nil
}
