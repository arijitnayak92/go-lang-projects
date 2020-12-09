CREATE TABLE "user" (
    "emailId" varchar PRIMARY KEY,
    "password" varchar NOT NULL,
    "loginIP" varchar NOT NULL,
    "isPasswordChanged" boolean NOT NULL
);
