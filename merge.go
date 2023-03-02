package schema2json

import "encoding/json"

func MergeSchemas(base, merge *Schema) (*Schema, error) {
	merged := new(Schema)

	jsonBase, err := json.Marshal(base)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBase, merged)
	if err != nil {
		return nil, err
	}
	jsonMerge, err := json.Marshal(merge)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonMerge, merged)
	if err != nil {
		return nil, err
	}

	return merged, nil
}

// func MergeSchemas(base, merge *Schema) (*Schema, error) {
// 	merged := map[string]interface{}{}
// 	result := new(Schema)

// 	jsonBase, err := json.Marshal(base)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(jsonBase, &merged)
// 	if err != nil {
// 		return nil, err
// 	}

// 	jsonMerge, err := json.Marshal(merge)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(jsonMerge, &merged)
// 	if err != nil {
// 		return nil, err
// 	}

// 	jsonResult, err := json.Marshal(merged)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(jsonResult, result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }
