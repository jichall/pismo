# The challenge

One is set to implement an interface to a financial service which has the
following actions:

- Create an account
- Create an transaction

# Prepare

To use the application you'll first have to set up the database environment. The
project defaults to PostgreSQL and as such you will need to have this database
installed.

I prefer to set up an entire database infrastructure that contains the tables
for the application and for that I create an user and a database.

```
# Log yourself in using the postgres user and open the PostgreSQL interface
$ sudo su postgres
$ psql
# Create a role and a database
CREATE ROLE pismo WITH PASSWORD 'a-very-strong-password';
ALTER ROLE pismo WITH LOGIN CREATEDB CREATEROLE;
CREATE DATABASE pismo;
```

Depending on your PostgreSQL system configuration you might not be able to
execute the following command. If so, alter the login method from **peer** to
**md5** in the file `etc/postgresql/10/main/pg_hba.conf`

```
$ cat *.sql | psql -U pismo -d pismo -W -1
```

# Build

To use the apaplication natively you can just build it with the following
command.

```
go build -o pismo.out src/*.go
```

After that you simply execute the binary and voilá.


# Docker

