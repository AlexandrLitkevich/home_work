package internalhttp

import (
	"net/http"

	"github.com/pressly/goose/v3"
)

func migrationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Выполните миграцию

		
		if err := goose.RunContext(db, "migrations", "up"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Вызовите следующий обработчик
		next.ServeHTTP(w, r)
	})
}
