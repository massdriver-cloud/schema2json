package schema2json

import (
	"fmt"
	"math/rand"
)

func generateArray(property *Schema) ([]interface{}, error) {
	if property.Const != nil {
		value, ok := property.Const.([]interface{})
		if !ok {
			return nil, fmt.Errorf("%s: unable to convert const %v to array", property.Name, property.Const)
		}
		return value, nil
	}

	result := []interface{}{}

	minItems := 0
	if property.MinItems != nil {
		minItems = int(*property.MinItems)
	}

	maxItems := 3
	if property.MaxItems != nil {
		maxItems = int(*property.MaxItems)
	}

	numItems := minItems + rand.Intn(maxItems-minItems)
	for idx := 0; idx < numItems; idx++ {
		element, err := generateValue(property.Items)
		if err != nil {
			return nil, err
		}
		result = append(result, element)
	}

	return result, nil
}
