package helpers

import (
	"fmt"
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetBlockInfoIDs(record [][]string) ([]int, error) {
	var index []int

	for i, row := range record {
		if len(row) == 0 {
			continue
		}

		if strings.Contains(strings.ToLower(strings.TrimSpace(row[0])), "block code") {
			index = append(index, i)
		}
	}

	return index, nil
}

func GetBlockInfo(record [][]string, index []int) ([]data.Block, error) {
	var blocks []data.Block
	blockInfo, err := GetCodeName(record, index)
	if err != nil {
		return nil, fmt.Errorf("error getting the code and name of block:%w", err)
	}

	for i, _ := range index {
		blocks = append(blocks, data.Block{
			BlockCode: blockInfo[i].Code,
			BlockName: blockInfo[i].Name,
		})
	}
	return blocks, nil

}
