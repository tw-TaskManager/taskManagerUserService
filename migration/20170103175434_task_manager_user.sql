
-- +goose Up
CREATE TABLE task_manager_user(
    "userName" text NOT NULL,
    "userId" serial primary key NOT NULL,
    "emailId" text NOT NULL,
    "password" text NOT NULL,
    unique("emailId")
);
-- SQL in section 'Up' is executed when this migration is applied

-- +goose Down
DROP TABLE task_manager_user;
-- SQL section 'Down' is executed when this migration is rolled back
