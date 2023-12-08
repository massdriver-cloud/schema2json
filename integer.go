package schema2json

import (
	"fmt"
	"math"
	"math/rand"
)

func generateInteger(property *Schema) (int, error) {
	if property.Default != nil {
		value, err := interfaceToInt(property.Default)
		if err != nil {
			return 0, fmt.Errorf("%s: unable to convert const %v to int", property.Name, property.Default)
		}
		return value, nil
	}

	if property.Enum != nil {
		idx := rand.Intn(len(property.Enum))
		value, err := interfaceToInt(property.Enum[idx])
		if err != nil {
			return 0, fmt.Errorf("%s: unable to convert const %v to int", property.Name, property.Enum[idx])
		}
		return value, nil
	}

	minimum := -9999999
	if property.Minimum != nil {
		minimum = int(*property.Minimum)
	}
	if property.ExclusiveMinimum != nil && *property.ExclusiveMinimum {
		minimum += 1
	}

	maximum := 99999999
	if property.Maximum != nil {
		maximum = int(*property.Maximum)
	}
	if property.ExclusiveMaximum != nil && *property.ExclusiveMaximum {
		maximum -= 1
	}

	if property.MultipleOf != nil {
		lowMultiplier := int(math.Floor(float64(minimum) / *property.MultipleOf))
		highMultiplier := int(math.Floor(float64(maximum) / *property.MultipleOf))
		multiplier := rand.Intn(highMultiplier-lowMultiplier) + lowMultiplier + 1
		return multiplier * int(*property.MultipleOf), nil
	}

	// adding 1 since Intn is exclusive at the top end
	value := rand.Intn(maximum-minimum+1) + minimum

	return value, nil
}

// // Golang assumes all numeric fields are float64 when deserializing anonymous JSON
func interfaceToInt(value interface{}) (int, error) {
	float, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("unable to convert %v to numeric field", value)
	}
	// need to tolerate a margin for float precision issues
	margin := 1e-15
	if _, frac := math.Modf(float); frac > margin || frac < -margin {
		ok = false
	}
	return int(float), nil
}
