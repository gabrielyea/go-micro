CREATE TABLE IF NOT EXISTS users(
   id serial NOT NULL,
   email VARCHAR (50) UNIQUE NOT NULL,
   first_name VARCHAR (50) NOT NULL,
   last_name VARCHAR (50) NOT NULL,
   password_hash VARCHAR NOT NULL,
   role INT
);