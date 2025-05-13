package middlware

import (
	"fmt"
	"net/http"

	"github.com/xadhithiyan/videon/service/auth"
	"github.com/xadhithiyan/videon/utils"
)

func AuthVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// the websocket must be removed from here
		if r.URL.Path == "/api/v1/login" || r.URL.Path == "/api/v1/register" {
			next.ServeHTTP(w, r)
			return
		}
		token, err := r.Cookie("token")
		if err != nil || token == nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not authenticated"))
			return
		}

		if _, ok := auth.AuthenticateJwt(token.Value); !ok {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not authenticated"))
			return
		}

		next.ServeHTTP(w, r)

	})
}
