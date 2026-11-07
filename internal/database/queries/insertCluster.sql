INSERT INTO clusters (
	block_code,
	cluster_code,
	cluster_name
)
VALUES (?,?,?)
ON CONFLICT DO NOTHING;
