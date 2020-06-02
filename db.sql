CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    password TEXT,
    level INTEGER NOT NULL
)

CREATE TABLE videos(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    descritpion TEXT,
    filepath TEXT NOT NULL
)
