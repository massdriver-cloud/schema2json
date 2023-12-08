package schema2json

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateJSON(schema *Schema) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	return generateValue(schema)
}

func generateValue(property *Schema) (interface{}, error) {
	if property.Const != nil {
		return property.Const, nil
	}

	// If the type isn't set at this point, we should assume its an 'object'
	if property.Type == nil {
		ty := "object"
		property.Type = &ty
	}

	if property.OneOf != nil {
		var err error
		property, err = selectOneOf(property)
		if err != nil {
			return nil, err
		}
	}

	switch *property.Type {
	case "array":
		return generateArray(property)
	case "boolean":
		return generateBoolean(property)
	case "integer":
		return generateInteger(property)
	case "number":
		return generateNumber(property)
	case "object", "":
		return generateObject(property)
	case "string":
		return generateString(property)
	default:
		return nil, fmt.Errorf("unsupported type: %s", *property.Type)
	}
}

func selectOneOf(property *Schema) (*Schema, error) {
	if len(property.OneOf) == 0 {
		return nil, fmt.Errorf("invalid schema: property %s contains 'oneOf' with no elements", property.Name)
	}
	return MergeSchemas(property, property.OneOf[rand.Intn(len(property.OneOf))])
}
