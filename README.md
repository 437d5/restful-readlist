# restful_readlist
The first RESTful API, with CRUD operations provided. It uses JSON to exchange data from client to server and vice-versa. It use PostgreSQL.
U need to setup database with 
```
   CREATE DATABASE test;
   CREATE TABLE readlist(
      id SERIAL PRIMARY KEY,
      done BOOL NOT NULL,
      author VARCHAR(100) NOT NULL,
      title VARCHAR(100) NOT NULL,
      year_published INT NOT NULL,
      rating INT CHECK (rating >= 0 AND rating <= 10)
)
```
After you did it u need to edit connString const in main.go file.
Enjoy.
