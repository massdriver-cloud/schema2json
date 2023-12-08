package schema2json

import (
	"fmt"
	"math/rand"

	"github.com/lucasjones/reggen"
)

func generateString(property *Schema) (string, error) {
	if property.Default != nil {
		value, ok := property.Default.(string)
		if !ok {
			return "", fmt.Errorf("%s: unable to convert default %v to string", property.Name, property.Default)
		}
		return value, nil
	}

	if property.Enum != nil {
		idx := rand.Intn(len(property.Enum))
		value, ok := property.Enum[idx].(string)
		if !ok {
			return "", fmt.Errorf("%s: unable to convert enum %v to string", property.Name, property.Enum[idx])
		}
		return value, nil
	}

	minLength := 4
	if property.MinLength != nil {
		minLength = *property.MinLength
	}
	maxLength := 20
	if property.MaxLength != nil {
		minLength = *property.MaxLength
	}

	pattern := fmt.Sprintf("^[a-zA-Z0-9!@#$^&?]{%d,%d}$", minLength, maxLength)
	if property.Pattern != nil {
		pattern = *property.Pattern
	}

	return reggen.Generate(pattern, maxLength)
}
