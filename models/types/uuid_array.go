package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("failed to parse UUIDArray: unsupport data type")
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	parts := strings.Split(str, ",")

	*a = make(UUIDArray, 0, len(parts))

	for _, s := range parts {
		s = strings.TrimSpace(strings.Trim(s, `"`))

		if s == "" {
			continue
		}

		u, err := uuid.Parse(s)
		if err != nil {
			return fmt.Errorf("Invalid UUID in Array : %v", err)
		}

		*a = append(*a, u)
	}

	return nil
}
