CREATE KEYSPACE IF NOT EXISTS key_gen_space
WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

CREATE TABLE IF NOT EXISTS key_gen_space.used_keys (
  value TEXT,
  create_date TIMESTAMP,
  PRIMARY KEY(value)
);

CREATE TABLE IF NOT EXISTS key_gen_space.unused_keys (
  value TEXT,
  create_date TIMESTAMP,
  PRIMARY KEY(value)
);
