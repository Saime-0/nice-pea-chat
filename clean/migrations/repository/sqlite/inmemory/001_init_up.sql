CREATE TABLE chats
(
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     creator_id INTEGER,
--     FOREIGN KEY (creator_id) REFERENCES users ON DELETE SET NULL
);