export POSTGRESQL_URL='postgres://postgres:admin@localhost:5432/meetup_dev?sslmode=disable'

Migration down
migrate -database ${POSTGRESQL_URL} -path graph/postgres/migrations down

Migration up
migrate -database ${POSTGRESQL_URL} -path graph/postgres/migrations up

Create table users
migrate create -ext sql -dir postgres/migrations create_users