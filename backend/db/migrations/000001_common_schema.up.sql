CREATE TABLE message_information (
  message_id BIGINT  AUTO_INCREMENT,
  title VARCHAR(255),
  user_id BIGINT,
  broker VARCHAR(255),
  created_at BIGINT,
  PRIMARY KEY(message_id)
);

CREATE TABLE users (
  user_id BIGINT NOT NULL,
  username VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  gender VARCHAR(255),
  dob BIGINT,
  created_at BIGINT,
  PRIMARY KEY (user_id)
);
