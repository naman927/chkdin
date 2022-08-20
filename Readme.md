Build a fully-fledged REST API that exposes GET, POST, DELETE and PUT endpoints that will subsequently allow you to perform the full range of CRUD operations, you can use any sql or nosql database. There should be provision to login and only Authenticated users should be able to do all the CRUD operations.

## NOTE: 
 - run this command to run a migration after doing `docker-compose up` `migrate -path db/migration -database "mysql://user:pass@tcp(localhost:3306)/demo" -verbose up`
