CREATE TABLE `message_information` (
  `message_id` integer AUTO_INCREMENT,
  `title` varchar(255),
  `user_id` integer,
  `tour_id` integer,
  `payment_id` integer,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(message_id)
);
