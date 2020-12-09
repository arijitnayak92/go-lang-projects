CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "productCategory"  (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    "description" varchar
);