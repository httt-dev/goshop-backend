-- +migrate Up
create table "users" (
"user_id" varchar(255) PRIMARY KEY,
"email" varchar(255) UNIQUE NOT NULL ,
"phone" varchar(10) UNIQUE NOT NULL,
"password" varchar(255)  NOT NULL ,
"address" varchar(255),
"full_name" varchar(150) ,
"avatar" text,
"role" varchar(25),
"created_at" TIMESTAMP NOT NULL ,
"updated_at" TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE users;
