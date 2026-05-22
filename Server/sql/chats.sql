CREATE TABLE IF NOT EXISTS chats (
    id SERIAL UNIQUE PRIMARY KEY NOT NULL,
    first_person_id INTEGER NOT NULL,
    second_person_id INTEGER NOT NULL,
    name TEXT,
    last_message TEXT,
    avatar_link TEXT,
    unread_count TEXT,
    updated_at TIMESTAMPTZ,
    created_at TEXT,

    CONSTRAINT  unique_chat UNIQUE (first_person_id, second_person_id)
)
