package graph

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalJSONScalar(j map[string]interface{}) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_ = json.NewEncoder(w).Encode(j)
	})
}

func UnmarshalJSONScalar(v interface{}) (map[string]interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		return val, nil
	default:
		return nil, fmt.Errorf("unexpected type for JSON scalar: %T", v)
	}
}
