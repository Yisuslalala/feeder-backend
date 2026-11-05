BEGIN;

CREATE TABLE IF NOT EXISTS feeder_details (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  feed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL DEFAULT 'member',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ NULL
);

-- Emails are case-insensitive, and a deleted user's email can be reused
CREATE UNIQUE INDEX IF NOT EXISTS uq_users_email
  ON users (lower(email))
  WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS houses (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(100),
  owner_user_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ NULL,

  CONSTRAINT fk_houses_owner
    FOREIGN KEY (owner_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS feeders (
  id BIGINT GENERATED ALWAYS AS IDENTITY,
  house_id BIGINT NOT NULL,
  mac_address CHAR(17) NOT NULL,
  name VARCHAR(100),
  pet_type VARCHAR(50),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ NULL,

  CONSTRAINT pk_feeders PRIMARY KEY (id),
  CONSTRAINT fk_feeders_house
    FOREIGN KEY (house_id)
    REFERENCES houses(id)
);

-- MAC addresses are case-insensitive, and a deleted feeder's MAC can be reused
CREATE UNIQUE INDEX IF NOT EXISTS uq_feeders_mac
  ON feeders (lower(mac_address))
  WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS house_users (
  house_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  role VARCHAR(20) NOT NULL DEFAULT 'member',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT pk_house_users
    PRIMARY KEY (house_id, user_id),

  CONSTRAINT ck_house_users_role
    CHECK (role IN ('owner', 'admin', 'member')),

  CONSTRAINT fk_house_users_house
    FOREIGN KEY (house_id)
    REFERENCES houses(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,

  CONSTRAINT fk_house_users_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

-- Postgres doesn't index foreign keys automatically, unlike MySQL
CREATE INDEX IF NOT EXISTS idx_houses_owner_user_id ON houses (owner_user_id);
CREATE INDEX IF NOT EXISTS idx_feeders_house_id ON feeders (house_id);
CREATE INDEX IF NOT EXISTS idx_house_users_user_id ON house_users (user_id);

COMMIT;
