CREATE TABLE message_information (
  message_id BIGINT  AUTO_INCREMENT,
  title VARCHAR(255),
  user_id BIGINT,
  broker VARCHAR(255),
  created_at BIGINT,
  PRIMARY KEY(message_id)
);