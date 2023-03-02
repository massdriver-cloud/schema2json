package schema2json

import (
	"fmt"
	"math/rand"
)

func generateBoolean(property *Schema) (bool, error) {
	if property.Const != nil {
		value, ok := property.Const.(bool)
		if !ok {
			return false, fmt.Errorf("%s: unable to convert const %v to bool", property.Name, property.Const)
		}
		return value, nil
	}

	return rand.Intn(2) == 1, nil
}
