package utility

import "fmt"

func EnsureKeys(requiredKeys []string, data map[string]any) error {
	for _, key := range requiredKeys {
		if _, ok := data[key]; !ok {
			return fmt.Errorf("the required field '%s' is missing", key)
		}
	}

	return nil
}
