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

# Console Output Example
2023/02/02 23:11:58 connect to http://localhost:8080/ for GraphQL playground
[cors] 2023/02/02 23:13:06 Handler: Actual request
[cors] 2023/02/02 23:13:06   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((email = 'zach@gmail.com')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
2023/02/02 23:13:06 [kbui-MBP/lrSvqi0wkJ-000001] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 233B in 135.600542ms
[cors] 2023/02/02 23:14:07 Handler: Actual request
[cors] 2023/02/02 23:14:07   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
SELECT "meetup"."id", "meetup"."name", "meetup"."description", "meetup"."user_id" FROM "meetups" AS "meetup" WHERE (id = '1') ORDER BY "meetup"."id" LIMIT 1
meetup:  1
2023/02/02 23:14:07 [kbui-MBP/lrSvqi0wkJ-000002] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 75B in 26.678042ms
[cors] 2023/02/02 23:16:01 Handler: Actual request
[cors] 2023/02/02 23:16:01   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
INSERT INTO "meetups" ("id", "name", "description", "user_id") VALUES (DEFAULT, 'Zach meetup', 'Zach description', '6') RETURNING *
2023/02/02 23:16:01 [kbui-MBP/lrSvqi0wkJ-000003] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 90B in 10.597959ms
[cors] 2023/02/02 23:16:18 Handler: Actual request
[cors] 2023/02/02 23:16:18   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
SELECT "meetup"."id", "meetup"."name", "meetup"."description", "meetup"."user_id" FROM "meetups" AS "meetup" WHERE (name ILIKE '%zach%') ORDER BY "id" LIMIT 5
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id in ('6'))) AND "user"."deleted_at" IS NULL
2023/02/02 23:16:18 [kbui-MBP/lrSvqi0wkJ-000004] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 145B in 10.966792ms
[cors] 2023/02/02 23:16:28 Handler: Actual request
[cors] 2023/02/02 23:16:28   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
SELECT "meetup"."id", "meetup"."name", "meetup"."description", "meetup"."user_id" FROM "meetups" AS "meetup" WHERE (id = '6') ORDER BY "meetup"."id" LIMIT 1
meetup:  
2023/02/02 23:16:28 [kbui-MBP/lrSvqi0wkJ-000005] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 84B in 1.244334ms
[cors] 2023/02/02 23:16:38 Handler: Actual request
[cors] 2023/02/02 23:16:38   Actual response added headers: map[Access-Control-Allow-Credentials:[true] Access-Control-Allow-Origin:[http://localhost:8080] Vary:[Origin]]
SELECT "user"."id", "user"."username", "user"."email", "user"."password", "user"."first_name", "user"."last_name", "user"."created_at", "user"."updated_at", "user"."deleted_at" FROM "users" AS "user" WHERE ((id = '6')) AND "user"."deleted_at" IS NULL ORDER BY "user"."id" LIMIT 1
SELECT "meetup"."id", "meetup"."name", "meetup"."description", "meetup"."user_id" FROM "meetups" AS "meetup" WHERE (id = '2') ORDER BY "meetup"."id" LIMIT 1
meetup:  2
DELETE FROM "meetups" AS "meetup" WHERE (id = '2')
2023/02/02 23:16:38 [kbui-MBP/lrSvqi0wkJ-000006] "POST http://localhost:8080/query HTTP/1.1" from [::1]:56003 - 200 30B in 6.598875ms