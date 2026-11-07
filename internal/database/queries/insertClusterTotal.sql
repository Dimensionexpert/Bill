INSERT INTO cluster_reports (
    scope, section, year, month, cluster_code,total_amount
)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(scope, section, year, month, cluster_code)
DO UPDATE SET
    total_amount = excluded.total_amount;
