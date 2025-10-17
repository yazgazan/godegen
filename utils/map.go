package utils

import (
	"errors"
	"fmt"
)

type Map struct{}

func (Map) New(values ...any) (map[string]any, error) {
	m := map[string]any{}

	if len(values)%2 != 0 {
		return m, errors.New("Map requires an even number of arguments")
	}

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return m, fmt.Errorf("Map keys must be strings (%d)", i)
		}

		m[key] = values[i+1]
	}

	return m, nil
}

func (Map) Set(key string, value any, m map[string]any) map[string]any {
	if m == nil {
		m = map[string]any{}
	}

	m[key] = value

	return m
}
