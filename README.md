# Bill

Bill is a Go-based parser for government school meal bill data.

## Features

- Parses CSV files converted from Excel bill reports.
- Extracts block, cluster, and school information.
- Calculates and validates cluster-wise and block-wise totals.
- Stores parsed data in SQLite.

## Tech Stack

- Go
- SQLite

## Validation

- **`ValidateTotalAmountsCluster()`**: Calculates the total meals served and total amount across schools within each cluster. Ensures that the parsed school data matches the cluster totals provided in the source.
- **`ValidateBlockTotals()`**: Calculates the total meals served and total amount across clusters within a block. Ensures that the parsed cluster totals match the block totals provided in the source.

## TODO

- [ ] Add input validation at the beginning of the transaction in `InsertIntoDB()`.
  - Validate required input slices.
  - Ensure related slices have matching lengths before indexing.
  - Prevent empty or invalid data from reaching database insertion.
- [ ] Add bank_info table if needed.
- [ ] Implement bill template generation. 
- [ ] Improve database integration and error handling.
- [ ] Consider Wails integration for a desktop application, depending on project requirements.

## Status

Currently under development. Database integration is in progress, and bill template generation is yet to be implemented.
