package data

// Data holds two values extracted from a CSV record.
//
// First and Second represent the order in which values appear in the
// parsed data, rather than specific field meanings. Their meanings depend
// on the parsing function that populates them.
//
// For example:
//   - GetCodeName: First is the code and Second is the name.
//   - GetServedAmount: First is the total served and Second is the total amount.
//
// This shared structure avoids duplicating parsing logic for different
// types of data that follow the same two-value pattern.

type Data struct {
	First  string
	Second string
}


