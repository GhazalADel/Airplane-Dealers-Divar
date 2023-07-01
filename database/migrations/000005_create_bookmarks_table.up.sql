CREATE TABLE IF NOT EXISTS bookmarks (
  user_id INTEGER, 
  ad_id INTEGER,
  PRIMARY KEY (user_id, ad_id),
  FOREIGN KEY user_id REFERENCES users(id),
  FOREIGN KEY ad_id REFERENCES ads(id)
);