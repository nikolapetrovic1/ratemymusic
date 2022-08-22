CREATE TABLE comments (
  id SERIAL PRIMARY KEY NOT NULL,
  comment TEXT NOT NULL,
  user_id integer NOT NULL,
  review_id integer NOT NULL,
  reports integer NOT NULL
)