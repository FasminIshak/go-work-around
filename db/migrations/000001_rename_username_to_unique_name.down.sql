-- Add old column back
ALTER TABLE users ADD COLUMN username text;

-- Copy data back
UPDATE users SET username = unique_name WHERE username IS NULL;

-- Make it not null
ALTER TABLE users ALTER COLUMN username SET NOT NULL;

-- Add unique constraint back
ALTER TABLE users ADD CONSTRAINT users_username_key UNIQUE (username);

-- Drop new column
ALTER TABLE users DROP COLUMN unique_name;
