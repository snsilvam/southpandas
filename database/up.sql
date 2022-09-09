DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS userClient;

CREATE TABLE userClient (
    id VARCHAR(32) PRIMARY KEY,
    premium VARCHAR(255) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT userClient_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

DROP TABLE IF EXISTS usersouthpandas;

CREATE TABLE usersouthpandas (
    id VARCHAR(32) PRIMARY KEY,
    type VARCHAR(32) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT usersouthpandas_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

DROP TABLE IF EXISTS workersofclient;

CREATE TABLE workersofclient (
    id VARCHAR(32) PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    userclient_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT workersofclient_userclient FOREIGN KEY (userclient_id) REFERENCES userclient (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

DROP TABLE IF EXISTS userexternalworker;

CREATE TABLE userexternalworker (
    id VARCHAR(32) PRIMARY KEY,
    contracttype VARCHAR(32) NOT NULL,
    workexperience VARCHAR(32) NOT NULL,
    workremote VARCHAR(255) NOT NULL,
    willingnesstravel VARCHAR(255) NOT NULL,
    currentsalary VARCHAR(32) NOT NULL,
    expectedsalary VARCHAR(32) NOT NULL,
    possibilityofrotation VARCHAR(255) NOT NULL,
    profilelinkedln VARCHAR(32) NOT NULL,
    workarea VARCHAR(32) NOT NULL,
    descriptionworkarea VARCHAR(255) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT userexternalworker_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);







