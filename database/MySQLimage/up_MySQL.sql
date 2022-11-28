CREATE DATABASE dummyusers;

USE dummyusers;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    -- user_id INT NOT NULL AUTO_INCREMENT,
    -- Library github.com/segmentio/ksuid
    -- user_id VARCHAR(32) PRIMARY KEY,
    id VARCHAR(32) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    -- firstName VARCHAR(100) NOT NULL,
    -- surname VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- considerar ANTES user_id ahora id 