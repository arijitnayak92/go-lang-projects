
CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "password" varchar,
  "firstName" varchar,
  "lastName" varchar,
  "role" varchar,
  "createdAt" timestamp,
   "cartId" int
);

CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "price" int,
  "imageId" int,
  "description" varchar,
  "createdAt" timestamp,
  "updatedAt" timestamp DEFAULT (now())
);
