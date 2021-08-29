CREATE KEYSPACE IF NOT EXISTS core_space
WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

CREATE TABLE IF NOT EXISTS core_space.urls (
  original_url TEXT,
  short_url TEXT,
  create_date TIMESTAMP,
  PRIMARY KEY(short_url)
);
