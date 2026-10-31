package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func Open(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(":( can't open the file-%s : %w", filepath.Base(path), err)
	}
	defer file.Close()

	record, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf(":( error reading csv file-%s : %w", path, err)
	}

	return record, nil
}
