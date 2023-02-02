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

# Run application
go run ./graph/server/server.go

# Regenerate Tables
go generate ./...

# GitHub libraries
go get github.com/go-chi/chi
go get github.com/rs/cors

# GraphQL Queries / Headers
query User {
  user(id: "1") {
    id
    username
    email
    meetups {
      id
      name
    }
  }
}

query GetMeetups {
  meetups(filter: {name: "ir"}, limit: 5, offset: 0) {
    name
    description
    user_id {
      id
      username
      email
    }
  }
}

mutation CreateMeetup {
  createMeetup(input: {name: "Lam meetup", description: "Lam description"}) {
    id
    name
    description
  }
}

mutation UpdateMeetup {
  updateMeetup(id: "1", input: {name: "a new name"}) {
    id
    name
  }
}

mutation DeleteMeetup {
  deleteMeetup(id: "1")
}

mutation RegisterUser {
  register(
    input: {email: "zach@gmail.com", username: "zachary", firstName: "Zachary", lastName: "Larouche", password: "tester123", confirmPassword: "tester123"}
  ) {
    authToken {
      accessToken
      expiredAt
    }
    user {
      id
      email
      username
    }
  }
}

mutation LoginUser {
  login(input: {email: "zach@gmail.com", password: "tester123"}) {
    authToken {
      accessToken
    }
    user {
      id
    }
  }
}

-Header
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU5MTEyNDYsImp0aSI6IjUiLCJpYXQiOjE2NzUzMDY0NDYsImlzcyI6Im1lZXRtZXVwIn0.Gb0InRkeRBwY5VAN_esVYQA3ihgow-4rw2d7otxNnqM"
}