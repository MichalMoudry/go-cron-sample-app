CREATE TABLE job_results (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(100) NOT NULL,
    started TIMESTAMP NOT NULL,
    finished TIMESTAMP NOT NULL
);