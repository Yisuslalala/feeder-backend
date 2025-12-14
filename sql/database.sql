CREATE DATABASE IF NOT EXISTS feeder;

USE feeder;

CREATE TABLE IF NOT EXISTS  feeder_details(
  id INT AUTO_INCREMENT PRIMARY KEY,
  feed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(50) DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL

  CONSTRAINT uq_users_email UNIQUE(email)
);

CREATE IF NOT EXISTS houses (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100),
  owner_user_id BIGINT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL

  CONSTRAINT fk_houses_owner
    FOREIGN KEY (owner_user_id) REFFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS house_users (
  house_id BIGINT UNSIGNED NOT NULL,
  user_id BIGINT UNSIGNED NOT NULL,
  role ENUM('owner', 'admin', 'member') NOT NULL DEFAULT 'member',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT pk_house_users
    PRIMARY KEY (house_id, user_id),

  CONSTRAINT fk_house_users_house
    FOREIGN KEY (house_id)
    REFERENCES houses(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE

  CONSTRAINT fk_house_users_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS feeders (
  id BIGINT UNSIGNED AUTO_INCREMENT,
  house_id BIGINT UNSIGNED NOT NULL,
  mac_address CHAR(17) NOT NULL,
  name VARCHAR(100),
  pet_type VARCHAR(50),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL

  CONSTRAINT pk_feeders PRIMARY KEY (id),
  CONSTRAINT fk_house fk_feeders_house
    FOREIGN KEY (house_idd) REFERENCES hosues(id),
  CONSTRAINT uq_feeders_mac UNIQUE (mac_address)

) ENGINE=InnoDB;

