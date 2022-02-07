# hashing-password
An API for hashing password in PostgreSQL with Golang

### Using PostgreSQL to store Encrypted string (can be passwords ideally) using Salt secret

- How it works ?
  - Just executing the Docker Compose file after adding the proper Env Vars.
  - Login to Database to create table like below and for sure we can use something like Gorm:

```
create table hashed_password(
	password varchar(100) not null
);

```
- Test it :)

```
curl -d "password123" -X POST http://localhost:8080/password
```

- EndPoints
 - `localhost:8080/health` => Health status with database
 - `http://localhost:8080/password` => POST endpoint

### Dependencies:
- [Gorilla Mux](github.com/gorilla/mux)
- [Postgres Driver](github.com/lib/pq)
- [Logrus Package](github.com/sirupsen/logrus)
- [Crypto Package](golang.org/x/crypto)