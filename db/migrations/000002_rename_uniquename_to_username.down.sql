-- Add new column
ALTER TABLE users ADD COLUMN IF NOT EXISTS unique_name text;

-- Copy data from old column
UPDATE users SET unique_name = username WHERE unique_name IS NULL;

-- Make it not null
ALTER TABLE users ALTER COLUMN unique_name SET NOT NULL;

-- Add unique constraint
ALTER TABLE users ADD CONSTRAINT users_unique_name_key UNIQUE (unique_name);

-- Drop old column
ALTER TABLE users DROP COLUMN username;
