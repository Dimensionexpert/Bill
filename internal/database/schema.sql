PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS blocks (
    block_code TEXT PRIMARY KEY,
    block_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS clusters (
    block_code TEXT NOT NULL,
    cluster_code TEXT PRIMARY KEY,
    cluster_name TEXT NOT NULL,

    FOREIGN KEY (block_code)
        REFERENCES blocks(block_code)
);

CREATE TABLE IF NOT EXISTS schools (
    cluster_code TEXT NOT NULL,
    school_code TEXT PRIMARY KEY,
    school_name TEXT NOT NULL,

    FOREIGN KEY (cluster_code)
        REFERENCES clusters(cluster_code)
);

CREATE TABLE IF NOT EXISTS block_reports (
    scope TEXT NOT NULL
        CHECK (scope IN ('STATE', 'CENTRAL')),
    section TEXT NOT NULL
        CHECK (section IN ('PRIMARY', 'UPPER_PRIMARY')),
    year INTEGER NOT NULL
        CHECK (year > 0),
    month INTEGER NOT NULL
        CHECK (month BETWEEN 1 AND 12),
    block_code TEXT NOT NULL,
    total_served INTEGER NOT NULL
        CHECK (total_served >= 0),
    total_amount INTEGER NOT NULL
        CHECK (total_amount >= 0),

    PRIMARY KEY (scope, section, year, month, block_code),

    FOREIGN KEY (block_code)
        REFERENCES blocks(block_code)
);

CREATE TABLE IF NOT EXISTS cluster_reports (
    scope TEXT NOT NULL
        CHECK (scope IN ('STATE', 'CENTRAL')),
    section TEXT NOT NULL
        CHECK (section IN ('PRIMARY', 'UPPER_PRIMARY')),
    year INTEGER NOT NULL
        CHECK (year > 0),
    month INTEGER NOT NULL
        CHECK (month BETWEEN 1 AND 12),
    cluster_code TEXT NOT NULL,
    total_served INTEGER NOT NULL
        CHECK (total_served >= 0),
    total_amount INTEGER NOT NULL
        CHECK (total_amount >= 0),

    PRIMARY KEY (scope, section, year, month, cluster_code),

    FOREIGN KEY (cluster_code)
        REFERENCES clusters(cluster_code)
);

CREATE TABLE IF NOT EXISTS school_reports (
    scope TEXT NOT NULL
        CHECK (scope IN ('STATE', 'CENTRAL')),
    section TEXT NOT NULL
        CHECK (section IN ('PRIMARY', 'UPPER_PRIMARY')),
    year INTEGER NOT NULL
        CHECK (year > 0),
    month INTEGER NOT NULL
        CHECK (month BETWEEN 1 AND 12),
    school_code TEXT NOT NULL,
    served INTEGER NOT NULL
        CHECK (served >= 0),
    amount INTEGER NOT NULL
        CHECK (amount >= 0),

    PRIMARY KEY (scope, section, year, month, school_code),

    FOREIGN KEY (school_code)
        REFERENCES schools(school_code)
);

CREATE INDEX IF NOT EXISTS idx_clusters_block_code
    ON clusters(block_code);

CREATE INDEX IF NOT EXISTS idx_schools_cluster_code
    ON schools(cluster_code);
