ALTER TABLE
posts

ADD 
CONSTRAINT fkr_user FOREIGN KEY (user_id) REFERENCES users(id)