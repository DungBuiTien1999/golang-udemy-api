package main

// Auth protects roots which needs authenticated
// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if !helpers.IsAuthenticated(r) {
// 			session.Put(r.Context(), "error", "log in first")
// 			http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
