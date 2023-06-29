CREATE TABLE IF NOT EXISTS bookmarks (
  user_id integer REFERENCES users(id), 
  ad_id integer REFERENCES ads(id),
  PRIMARY KEY (user_id, ad_id)
);