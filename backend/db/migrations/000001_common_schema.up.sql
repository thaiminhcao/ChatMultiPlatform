CREATE TABLE message_information (
  message_id SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(255),
  user_id INT,
  broker VARCHAR(255),
  created_at BIGINT
);

CREATE TABLE users (
  user_id INT NOT NULL,
  username VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  gender VARCHAR(255),
  dob BIGINT,
  created_at BIGINT,
  PRIMARY KEY (user_id)
);
