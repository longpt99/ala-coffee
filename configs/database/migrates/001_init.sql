-- Create ext support using UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create common status enum
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
    CREATE TYPE status_enum AS ENUM ('active', 'inactive');
  END IF;
END $$;

-- Create gender enum
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
    CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');
  END IF;
END $$;