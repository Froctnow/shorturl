package mapper

import (
	"fmt"
	"strings"
)

func (m *mapper) URLIDs(urls *[]string) map[string]any {
	var IDs []string

	for _, url := range *urls {
		IDs = append(IDs, fmt.Sprintf("'%s'", url))
	}

	return map[string]any{
		"IDs": strings.Join(IDs, ","),
	}
}
