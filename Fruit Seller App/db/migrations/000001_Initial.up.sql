CREATE TABLE "users" (
    "email" varchar PRIMARY KEY,
    "password" varchar NOT NULL,
    "firstName" varchar NOT NULL,
    "lastName" varchar NOT NULL,
    "role"  varchar NOT NULL,
    "cartId" varchar NOT NULL,
);

CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "price" int,
  "imageId" varchar,
  "description" varchar,
  "createdAt" timestamp,
  "updatedAt" timestamp DEFAULT (now())
);

