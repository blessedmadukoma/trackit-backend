## Golang Backend Template

This is a template guide on how my Go backend projects are structured.

Features Covered:
1. user authentication i.e. registration and login using email with access and refresh tokens
2. session management

How to run:
1. Clone the repository.
2. Change the name of the `go.mod` file to a mod name of your choice, and replace in all locations.
3. Install all the packages: `go mod tidy`.
4. Set the env configuration.
5. Run `make sqlc` to update db queries and models.
6. First set up your database, the run `make migrate` to perform a migration.
7. Run `make test` to make sure all tests pass.
8. Spin up the project: `go run main.go` or for `air` users: `air`.