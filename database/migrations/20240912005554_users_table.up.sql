CREATE TABLE users (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL,
  PRIMARY KEY (id),
  UNIQUE KEY (email),
  KEY idx_users_created_at (created_at),
  KEY idx_users_updated_at (updated_at)
);
