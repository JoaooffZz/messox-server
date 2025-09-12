CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    password VARCHAR(250) NOT NULL,
    profile TEXT NOT NULL,
    bio TEXT NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS contacts (
    user_id INT NOT NULL,
    contact_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS chats (
    chat_id BIGINT NOT NULL,
    date DATE NOT NULL,
    chat JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS inbox_messages (
    sender_id INT NOT NULL,
    addressee_id INT NOT NULL,
    message TEXT NOT NULL,
    created_at TIME NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (addressee_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS inbox_requests (
    sender_id INT NOT NULL,
    addressee_id INT NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (addressee_id) REFERENCES users(id) ON DELETE CASCADE
);