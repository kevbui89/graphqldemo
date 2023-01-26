package graphql

import (
	"context"
	"net/http"
	"time"

	"com.example/graphql/graph/model"
	"github.com/go-pg/pg/v10"
)

const userLoaderKey = "userloader"

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				u := make(map[string]*model.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
