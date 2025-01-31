CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,  
  name TEXT,
  nik TEXT,
  phone TEXT UNIQUE,	
  number TEXT UNIQUE,
  balance BIGINT,
  created_at DATE
);

CREATE TABLE account_histories (
  id SERIAL PRIMARY KEY,
  account_id INT NOT NULL,
  type TEXT,
  nominal BIGINT,
  current_balance BIGINT,
  new_balance BIGINT,
  created_at DATE,
  FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
)