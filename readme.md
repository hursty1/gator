helpful commands:

psql -U postgres -h localhost -d gator

\c gator

\dt shows tables

sqlc generate


goose postgres postgres://postgres:postgres@localhost:5432/gator down

goose postgres postgres://postgres:postgres@localhost:5432/gator up