package helpers

import (
	"fmt"
	"strings"
)

func GetSectionScope(record [][]string) (section, scope string, err error) {
	if len(record[3]) < 4 {
		return "", "", fmt.Errorf(":( record[3]'s lenght is shorter than expected :%w", err)
	}

	switch strings.ToLower(strings.TrimSpace(record[3][2])) {
	case "class 1 to 5":
		section = "PRIMARY"
	case "class 6 to 8":
		section = "UPPER_PRIMARY"
	default:
		return "", "", fmt.Errorf("invalid section value: %q", record[3][2])
	}

	scpe := strings.ToLower(strings.TrimSpace(record[3][4]))
	switch {
	case strings.Contains(scpe, "central rate"):
		scope = "CENTRAL"
	case strings.Contains(scpe, "state rate"):
		scope = "STATE"
	default:
		return "", "", fmt.Errorf("invalid scope value: %q", record[3][4])
	}

	return section, scope, nil

}
