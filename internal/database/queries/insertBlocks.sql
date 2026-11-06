INSERT INTO blocks (
	block_code,
	block_name
)
VALUES (?,?)
ON CONFLICT DO NOTHING;
