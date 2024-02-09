CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    encrypted_password TEXT NOT NULL
);
CREATE TABLE activities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    DESCRIPTION TEXT NOT NULL,
    datetime TIMESTAMP
);
CREATE TABLE activities_registration (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    activity_id INT REFERENCES activities (id) ON UPDATE CASCADE ON DELETE CASCADE,
    UNIQUE (user_id, activity_id)
);
