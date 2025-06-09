package graph

import (
	"encoding/json"
	"fmt"
	db "goCrud/prisma/db/prisma-client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *db.PrismaClient
}

func MarshalJSON(v interface{}) (interface{}, error) {
	return v, nil
}

func UnmarshalJSON(v interface{}) (map[string]interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		return val, nil
	case string:
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(val), &m); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unexpected type for JSON: %T", v)
	}
}
