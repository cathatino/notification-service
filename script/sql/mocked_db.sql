`mocked user table`
CREATE TABLE mocked_user_tab (
	user_id serial PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	ctime INT NOT NULL
);

`client table`
CREATE TABLE client_tab (
    client_id serial PRIMARY KEY,
    client_category INT NOT NULL,
    client_key VARCHAR(64) NOT NULL,
    description VARCHAR(256) NOT NULL,
    ctime INT NOT NULL,
    mtime INT NOT NULL
);

