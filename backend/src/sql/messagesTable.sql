create table messages(
    id SERIAL PRIMARY KEY NOT NULL,
    text TEXT NOT NULL,
    created_at DATE NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    is_changed BOOLEAN NOT NULL
);