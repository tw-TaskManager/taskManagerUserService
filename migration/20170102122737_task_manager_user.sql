
-- +goose Up
CREATE TABLE task_manager_user(
    "userName" text NOT NULL,
    "emailId" text primary key NOT NULL,
    "password" text NOT NULL
);
-- SQL in section 'Up' is executed when this migration is applied

-- +goose Down
DROP TABLE task_manager_user;
-- SQL section 'Down' is executed when this migration is rolled back
