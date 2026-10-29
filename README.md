# Bill

Bill is a Go-based parser for government school meal bill data.

## Features

- Parses CSV files converted from Excel bill reports.
- Extracts block, cluster, and school information.
- Calculates and validates cluster-wise and block-wise totals.

## Tech Stack

- Go
- SQLite (planned)

## Status

Currently under developement. Database integration and template generation is yet to add.
Wails integretion might come. dependening upon the circumstences ;).

## validation
- ValidateTotalAmountsCluster() this function, calculates the served total and amount total of schools from that cluster from source. Ensuring that the parsed data is accurate, validating parse data of both school and cluster totals ( from source)
- ValidateBlockTotals() this function, calculates the served total and amount total of clusters from source. Ensuring that the parsed data is accurate, validating both cluster totals and block total.
