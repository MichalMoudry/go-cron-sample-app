CREATE TABLE job_results (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(100) NOT NULL,
    started VARCHAR(50) NOT NULL,
    finished VARCHAR(50) NOT NULL
);