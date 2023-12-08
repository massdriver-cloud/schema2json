package schema2json

import (
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
			dep, err := generateValue(schema)
			if err != nil {
				return result, err
			}
			maps.Copy(result, dep.(map[string]interface{}))

		}
	}

	return result, nil
}
