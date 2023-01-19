-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE comic (
  id SERIAL NOT NULL PRIMARY KEY,
  title VARCHAR(256), 
  description VARCHAR(256), 
  image_url VARCHAR(256), 
  release_year INT,
  genre VARCHAR(256), 
  type VARCHAR(256), 
  status BOOLEAN, 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE rating (
  id SERIAL NOT NULL PRIMARY KEY,
  comic_id INT, 
  user_id INT, 
  rate INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE comments (
  id SERIAL NOT NULL PRIMARY KEY,
  comic_id INT, 
  user_id INT, 
  reference_id INT,
  comment VARCHAR(256),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE chapters (
  id SERIAL NOT NULL PRIMARY KEY,
  comic_id INT, 
  chapter_number INT, 
  image_url VARCHAR(256),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd
