INSERT INTO block_reports (
    scope, section, year, month, block_code,total_amount
)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(scope, section, year, month, block_code)
DO UPDATE SET
    total_amount = excluded.total_amount;
