CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,

    project_id INT,

    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),

    assignee_id INT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (project_id)
    REFERENCES projects(id),

    FOREIGN KEY (assignee_id)
    REFERENCES users(id)
);