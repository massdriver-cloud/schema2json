package schema2json

import (
	"fmt"
	"math/rand"
)

func generateBoolean(property *Schema) (bool, error) {
	if property.Default != nil {
		value, ok := property.Default.(bool)
		if !ok {
			return false, fmt.Errorf("%s: unable to convert const %v to bool", property.Name, property.Default)
		}
		return value, nil
	}

	return rand.Intn(2) == 1, nil
}
