-- here up is the up migration and down is the down migration 
-- migration scripts are the one that automatically runs on all the environments

-- +goose up
CREATE TABLE users;
-- +goose down
DROP TABLE users;
