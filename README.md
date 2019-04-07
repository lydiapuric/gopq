# PostgreSQL with Go
Connect to a PostgreSQL database using Go.
You can pass the database configuration either as Command-Line Flags or as Environment Variables

## How to run

### Build application
    go build

This will create a binary file with the name __gopq__ in current directory.

### Run application and pass the database configuration using Command-Line Flags
    ./gopq -h=localhost -p=5432 -db=dbname -u=postgres -pwd=dbpwd -ssl=require

### Run application and pass the database configuration using Environment Variables
    DB_HOST=localhost DB_PORT=5432 DB_NAME=dbname DB_USER=postgres DB_PWD=dbpwd DB_SSL=require ./gopq

   