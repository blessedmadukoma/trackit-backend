## TrackIT Backend

This is backend for the TrackIT project..

Features:
1. user authentication
2. session management
3. transactions
4. invoice and expense

How to run:
1. Clone the repository.
2. Change the name of the `go.mod` file to a mod name of your choice, and replace in all locations.
3. Install all the packages: `go mod tidy`.
4. Set the env configuration.
5. Run `make sqlc` to update db queries and models.
6. First set up your database, the run `make migrate` to perform a migration.
7. Run `make test` to make sure all tests pass.
8. Spin up the project: `go run main.go` or for `air` users: `air`.

Solution to my docker nightmares:
First: build docker (i.e. build the Dockerfile, not docker-compose.yaml): `docker build -t trackit:latest .`