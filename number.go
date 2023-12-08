package schema2json

import (
	"fmt"
	"math"
	"math/rand"
)

func generateNumber(property *Schema) (float64, error) {
	if property.Default != nil {
		value, ok := property.Default.(float64)
		if !ok {
			return 0, fmt.Errorf("%s: unable to convert default %v to float64", property.Name, property.Default)
		}
		return value, nil
	}

	if property.Enum != nil {
		idx := rand.Intn(len(property.Enum))
		value, ok := property.Enum[idx].(float64)
		if !ok {
			return 0, fmt.Errorf("%s: unable to convert enum %v to float64", property.Name, property.Enum[idx])
		}
		return value, nil
	}

	minimum := -999999999999.9
	if property.Minimum != nil {
		minimum = *property.Minimum
	}
	if property.ExclusiveMinimum != nil && *property.ExclusiveMinimum {
		minimum += 1
	}

	maximum := 999999999999.9
	if property.Maximum != nil {
		maximum = *property.Maximum
		if property.ExclusiveMinimum != nil && *property.ExclusiveMinimum {
			maximum -= 1
		}
	}

	if property.MultipleOf != nil {
		lowMultiplier := int(math.Floor(minimum / *property.MultipleOf))
		highMultiplier := int(math.Floor(maximum / *property.MultipleOf))
		multiplier := rand.Intn(highMultiplier-lowMultiplier) + lowMultiplier + 1
		return float64(multiplier) * *property.MultipleOf, nil
	}

	value := minimum + rand.Float64()*(maximum-minimum)

	return value, nil
}
