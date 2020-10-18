package controllers

import "net/http"

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var header = req.Header.Get("x-access-token")

		header = strings.TrimSpace(header)

		if header == "" {
			res.WriteHeader(http.StatusForbidden)
			json.NewEncoder(res).Encode(Exception{
				Message: "Missing auth token"
			})
			return
		}

		tk = &models.LoginToken{}

		_, err := jwt.ParseWithClaims(header, tk, funct(token *jwt.Token) (interface{}, err){
			return []byte("secret"), nil
		})

		if err != nil {
			res.WriteHeader(http.StatusForbidden)
			json.NewEncoder(res).Encode(Exception{
				Message: err.Error()
			})
			return
		}
		
		ctx := context.WithValue(req.Context(), "user", tk)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
