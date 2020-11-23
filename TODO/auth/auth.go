package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		err := domain.UserMethodMux.TokenValid(req)
		fmt.Println("in auth")
		fmt.Println(err)
		if err != nil {
			fmt.Println("Got error")
			res.WriteHeader(401)
			jsonVaue, _ := json.Marshal(err)
			res.Write(jsonVaue)
			return
		}
		// accessDetails, _ := domain.UserMethodMux.ExtractTokenMetadata(req)
		next.ServeHTTP(res, req)
	})
}
