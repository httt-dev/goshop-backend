-- +migrate Up
create  table  "categories"(
    "cate_id" varchar (255) PRIMARY KEY,
    "cate_name" varchar (255) UNIQUE NOT NULL,
    "cate_image" varchar (255),
    "created_at" timestamp not null,
    "updated_at" timestamp not null
);

-- +migrate Down
Drop table categories;