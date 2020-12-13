
CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "password" varchar,
  "firstname" varchar,
  "lastname" varchar,
  "role" varchar, 
  "createdat" timestamp,
  "updatedat" timestamp default current_timestamp,
   "cartid" int
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
