CREATE TABLE IF NOT EXISTS bookmarks (
  user_id SERIAL,
  ads_id SERIAL,
  PRIMARY KEY (user_id, ads_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (ads_id) REFERENCES ads (id)
);
-- Create the join table for the many-to-many relationship
CREATE TABLE IF NOT EXISTS user_bookmark (
  user_id SERIAL,
  bookmark_user_id SERIAL,
  bookmark_ads_id SERIAL,
  PRIMARY KEY (user_id, bookmark_user_id, bookmark_ads_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (bookmark_user_id, bookmark_ads_id) REFERENCES bookmarks (user_id, ads_id)
);