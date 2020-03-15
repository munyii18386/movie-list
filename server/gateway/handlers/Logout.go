package handlers

import(
	"net/http"
	"fmt"
)

// LogoutHandler is used to log a user out of their account
func (ctx *Context) LogoutHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch{
		token := r.Header.Get("Authorization")
		fmt.Println("The logout token is: ", token)

		sid, err := EndSession(ctx.SessionKey, ctx.RedisDatabase, token)

		if err == nil {
			fmt.Println("this token and session has been successfully deleted", sid)
			w.WriteHeader(http.StatusOK)
		}
		

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	}

}