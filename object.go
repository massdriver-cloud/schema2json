package schema2json

import "fmt"

func generateObject(property *Schema) (map[string]interface{}, error) {
	if property.Const != nil {
		value, ok := property.Const.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("%s: unable to convert const %v to object", property.Name, property.Const)
		}
		return value, nil
	}

	result := map[string]interface{}{}
	var err error
	for name, property := range property.Properties {
		property.Name = name
		result[name], err = generateValue(property)
		if err != nil {
			return result, err
		}
	}

	if property.Dependencies != nil {
		for name, schema := range property.Dependencies {
			applyDependency(name, property, schema)
		}
	}

	return result, nil
}

func applyDependency(propertyName string, current, dependencySchema *Schema) error {
	return nil
}
