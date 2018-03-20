CREATE KEYSPACE IF NOT EXISTS posts WITH REPLICATION = {
  'class':'SimpleStrategy','replication_factor':1
};

CREATE TABLE IF NOT EXISTS posts.posts_flakeseries (
  bucket TIMESTAMP,
  flake_created TIMESTAMP,

  id TEXT,
  body TEXT,
  PRIMARY KEY ((bucket), flake_created, id)
);
