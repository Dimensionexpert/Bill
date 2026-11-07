INSERT INTO schools (
	cluster_code,
	school_code,
	school_name
)
VALUES (?,?,?)
ON CONFLICT DO NOTHING;
