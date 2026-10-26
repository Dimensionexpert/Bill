package helpers

import (
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetSchoolInfoIdx(record [][]string, blockCode string) ([]int, error) {
	var index []int
	for i, row := range record {
		if len(row) == 0 {
			continue
		}

		if strings.Contains(strings.ToLower(strings.TrimSpace(row[2])), blockCode) {
			index = append(index, i)
		}
	}

	return index, nil

}

func GetSchool(record [][]string, index []int) (data.School, error) {
	idx := index[0]

	return data.School{
		UDISECode:   strings.TrimSpace(record[idx][2]),
		SchoolName:  strings.TrimSpace(record[idx][4]),
		MealsServed: strings.TrimSpace(record[idx][6]),
		Amount:      strings.TrimSpace(record[idx][8]),
	}, nil
}

func GetSchools(record [][]string, indexes []int) ([]data.School, error) {
	var schools []data.School

	for _, idx := range indexes {
		school, err := GetSchool(record, []int{idx})
		if err != nil {
			return nil, err
		}

		schools = append(schools, school)
	}

	return schools, nil
}





//deprecated

// func GetUDISENameServedAmounts(record [][]string, index []int) ([]data.School, error) {
// 	var schoolInfo []data.School

// 	for _, idx := range index {
// 		udise := strings.TrimSpace(record[idx][2])
// 		schoolName := strings.TrimSpace(record[idx][4])
// 		mealServed := strings.TrimSpace(record[idx][6])
// 		amount := strings.TrimSpace(record[idx][8])

// 		schoolInfo = append(schoolInfo, data.School{
// 			UDISECode:   udise,
// 			SchoolName:  schoolName,
// 			MealsServed: mealServed,
// 			Amount:      amount,
// 		})

// 	}
// 	return schoolInfo, nil

// }
