CREATE TABLE message_information (
  message_id SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(255),
  user_id INT,
  broker VARCHAR(255),
  created_at BIGINT
);
