export POSTGRESQL_URL='postgres://postgres:admin@localhost:5432/meetup_dev?sslmode=disable'

# Migration down
migrate -database ${POSTGRESQL_URL} -path graph/postgres/migrations down

# Migration up
migrate -database ${POSTGRESQL_URL} -path graph/postgres/migrations up

# Create table users
migrate create -ext sql -dir postgres/migrations create_users

# Regenerate the schema.graphqls file
go run github.com/99designs/gqlgen

# Dataloden library
go get github.com/vektah/dataloaden
go run github.com/vektah/dataloaden UserLoader string '*com.example/graphql/graph/model.User'