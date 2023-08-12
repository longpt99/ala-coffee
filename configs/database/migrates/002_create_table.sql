-- Create a table called "categories"
CREATE TABLE IF NOT EXISTS categories (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at timestamp,
  PRIMARY KEY (id)
);

-- Create a table called "products"
CREATE TABLE IF NOT EXISTS products (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  categoryIds TEXT[] NOT NULL,
  status status_enum NOT NULL DEFAULT 'inactive',
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at timestamp,
  PRIMARY KEY (id)
);

-- Create a table called "product_categories"
CREATE TABLE IF NOT EXISTS product_categories (
  product_id UUID REFERENCES products(id),
  category_id UUID REFERENCES categories(id),
  PRIMARY KEY (product_id, category_id)
);

-- Create a table called "admins"
CREATE TABLE IF NOT EXISTS admins (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at timestamp,
  PRIMARY KEY (id)
);

-- Create a table called "users"
CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  dob VARCHAR(10),
  gender gender_enum,
  status status_enum DEFAULT 'active' NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  deleted_at timestamp,
  PRIMARY KEY (id)
);