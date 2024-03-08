package cors

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func CorsHandler() func(http.Handler) http.Handler {
	return handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)
}
