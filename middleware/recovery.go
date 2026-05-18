package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				fmt.Println("PANIC RECOVERED:", err)

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}